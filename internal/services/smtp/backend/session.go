package backend

import (
	"context"
	"io"

	"github.com/emersion/go-smtp"
	"github.com/jordan-wright/email"
	"github.com/lvlcn-t/DevSMTP/internal/models"
	"github.com/lvlcn-t/halog/pkg/logger"
)

type session struct {
	ctx     context.Context
	ctxCanc context.CancelFunc

	conn *smtp.Conn
	be   *backend
}

// NewSession creates a new SMTP session for each client connection.
// This function is called when a new client connection is made to the SMTP server.
func (b *backend) NewSession(conn *smtp.Conn) (smtp.Session, error) {
	ctx, cancel := context.WithCancel(b.ctx)
	return &session{
		ctx:     ctx,
		ctxCanc: cancel,
		be:      b,
		conn:    conn,
	}, nil
}

// AuthPlain handles the authentication step in the SMTP protocol.
// It is called when the client sends the HELO/EHLO command with authentication details.
// If authentication is not required (no user and password set in the config),
// it allows the session to proceed without authentication.
func (s *session) AuthPlain(username, password string) error {
	log := logger.FromContext(s.ctx)
	if s.be.cfg.User == "" && s.be.cfg.Password == "" {
		log.Info("Authentication skipped: no credentials configured")
		return nil
	}
	if username != s.be.cfg.User || password != s.be.cfg.Password {
		log.Warn("Authentication failed: invalid credentials")
		return smtp.ErrAuthFailed
	}
	log.Debug("Authentication successful")
	return nil
}

// Mail is called when the client sends the 'MAIL FROM' command.
// It sets up the sender's email address for the email transaction.
// Currently, this implementation does not perform any action with the sender's address.
func (s *session) Mail(from string, opts *smtp.MailOptions) error {
	return nil
}

// Rcpt is called for each 'RCPT TO' command issued by the client.
// It specifies a recipient of the email. This implementation does not perform
// any action with the recipient's address, but it could be used for validation or logging.
func (s *session) Rcpt(to string, opts *smtp.RcptOptions) error {
	return nil
}

// Data is called when the client sends the 'DATA' command and starts transmitting the actual email content.
// This method reads the email data, parses it, and saves it to the database.
// It returns an SMTP error if the email cannot be processed or saved.
func (s *session) Data(r io.Reader) error {
	log := logger.FromContext(s.ctx)

	e, err := email.NewEmailFromReader(r)
	if err != nil {
		log.Error("Failed to process email", "error", err)
		return &smtp.SMTPError{
			Code:         550,
			EnhancedCode: smtp.EnhancedCode{5, 0, 0},
			Message:      "Failed to process email",
		}
	}

	email := models.Email{
		Email: e,
	}

	if err := s.be.db.SaveEmail(email); err != nil {
		log.Error("Failed to save email to the database", "error", err)
		return &smtp.SMTPError{
			Code:         550,
			EnhancedCode: smtp.EnhancedCode{5, 0, 0},
			Message:      "Failed to process email",
		}
	}

	log.Info("Email processed and saved")
	return nil
}

// Reset is called to reset the state of the session.
// This function is used to clear any message-specific state and prepare for a new message.
// Currently, there's no specific state maintained in the session, so it does nothing.
func (s *session) Reset() {}

// Logout is called when the client sends the 'QUIT' command.
// It is used to clean up any session-specific resources and close the session.
func (s *session) Logout() error {
	defer s.ctxCanc()
	return nil
}
