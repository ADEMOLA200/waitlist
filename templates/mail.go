package templates

import (
	"bytes"
	"html/template"
	"time"
)

// CreditEmailData represents the data needed for rendering the credit notification email template.
type CreditEmailData struct {
	Recipient string
	Amount    int
	Timestamp string
}

// RenderCreditNotification renders the email body for credit notification using a template.
func RenderCreditNotification(email string, amount int) (string, error) {
    // Define email template HTML
    const emailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Your Credits</title>
</head>
<body>
    <p>Dear {{.Recipient}},</p>
    <p>You have received ${{.Amount}} credits!</p>
    <p>This email was sent at {{.Timestamp}}.</p>
    <p>Thank you.</p>
</body>
</html>`

    // Parse the email template
    tmpl, err := template.New("email").Parse(emailTemplate)
    if err != nil {
        return "", err
    }

    // Execute the template with user data
    var tpl bytes.Buffer
    data := CreditEmailData{
        Recipient: email,
        Amount:    amount,
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
    }
    if err := tmpl.Execute(&tpl, data); err != nil {
        return "", err
    }

    // Return the rendered HTML content as a string
    return tpl.String(), nil
}