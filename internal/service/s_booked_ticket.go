package service

import (
	"booking_service/internal/entity"
	"booking_service/internal/repository"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net/smtp"
)

type BookedTicketService struct {
	bookedTicketRepository repository.BookedTicket
}

func NewBookedTicketService(bookedTicketRepository *repository.Repository) *BookedTicketService {
	return &BookedTicketService{
		bookedTicketRepository: bookedTicketRepository,
	}
}

func (s *BookedTicketService) Create(ticket entity.BookedTicket) (int, error) {
	id, err := s.bookedTicketRepository.Create(ticket)
	if err != nil {
		return 0, err
	}

	// Отправкка уведомления
	err = sendHTMLMail("olardaniil@vk.com")
	if err != nil {
		return 0, err
	}

	return id, nil
}

// SendToRabbitMQ Функция для отправки структуры Income в RabbitMQ
func (s *BookedTicketService) SendToRabbitMQ(income entity.Income, queueName string, amqpURI string) error {
	// Подключение к RabbitMQ
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	defer conn.Close()

	// Создание канала
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	// Объявление очереди
	_, err = ch.QueueDeclare(
		queueName, // имя очереди
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	// Сериализация структуры Income в JSON
	body, err := json.Marshal(income)
	if err != nil {
		return fmt.Errorf("failed to marshal income to JSON: %w", err)
	}

	// Публикация сообщения
	err = ch.Publish(
		"",        // exchange
		queueName, // routing key (имя очереди)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	log.Printf("Successfully sent message to queue %s", queueName)
	return nil
}

func sendHTMLMail(to string) error {
	// Настройки SMTP-сервера
	smtpHost := "smtp.mail.ru"
	smtpPort := "587"

	senderEmail := "control.prices@mail.ru"
	password := "JFpkRRRUTe8wjSjn1e0V"

	subject := "Ваша заявка на автобронирование успешно выполнена!"
	htmlBody := "<!DOCTYPE html>\n<html lang=\"ru\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Уведомление</title>\n</head>\n<body>\n<div autoid=\"_rp_x\" class=\"_rp_T4\" id=\"Item.MessagePartBody\">\n    <div class=\"_rp_U4 ms-font-weight-regular ms-font-color-neutralDark rpHighlightAllClass rpHighlightBodyClass\"\n         id=\"Item.MessageUniqueBody\"\n         style=\"font-family: wf_segoe-ui_normal, &quot;Segoe UI&quot;, &quot;Segoe WP&quot;, Tahoma, Arial, sans-serif, serif, EmojiFont;\">\n        <div class=\"rps_fe50\">\n            <div>\n                <div lang=\"RU\" style=\"background-color:#f2f3f7; background-color:#f2f3f7\">\n                    <div style=\"display:none\"><img\n                            src=\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNkYAAAAAYAAjCB0C8AAAAASUVORK5CYII=\">\n                    </div>\n                    <div align=\"center\">\n                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                               style=\"width:100%; border-collapse:collapse\">\n                            <tbody>\n                            <tr style=\"background-color:#f2f3f7\">\n                                <td width=\"580px\" style=\"height:24px; line-height:21.8px; width:580px\">&nbsp;</td>\n                            </tr>\n                            </tbody>\n                        </table>\n                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                               style=\"width:100%; border-collapse:collapse\">\n                            <tbody>\n                            <tr style=\"background-color:#f2f3f7\">\n                                <td colspan=\"4\" style=\"width:24px; max-width:24px; min-width:24px\">&nbsp;</td>\n                                <td valign=\"bottom\" style=\"padding:0cm 0cm 0cm 0cm\">\n                                    <div align=\"center\">\n                                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"0\"\n                                               style=\"width:580px; border-collapse:collapse\">\n                                            <tbody>\n                                            <tr style=\"background-color:#444343; width:580px; height:280px; text-align:center\">\n                                                <td colspan=\"4\" style=\"padding:0\"><img\n                                                        src=\"https://i.postimg.cc/jSNLJGMF/Frame-432.png\"\n                                                        width=\"580\" height=\"280\" alt=\"Изображение\" style=\"margin:auto\">\n                                                </td>\n                                            </tr>\n                                            </tbody>\n                                        </table>\n                                    </div>\n                                </td>\n                                <td colspan=\"4\" style=\"width:24px; max-width:24px; min-width:24px\">&nbsp;</td>\n                            </tr>\n                            </tbody>\n                        </table>\n                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                               style=\"width:100%; border-collapse:collapse\">\n                            <tbody>\n                            <tr style=\"background-color:#f2f3f7\">\n                                <td colspan=\"4\" style=\"width:24px; max-width:24px; min-width:24px\">&nbsp;</td>\n                                <td valign=\"bottom\" style=\"padding:0cm 0cm 0cm 0cm\">\n                                    <div align=\"center\">\n                                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"580px\"\n                                               style=\"width:580px; border-collapse:collapse\">\n                                            <tbody>\n                                            <tr>\n                                                <td valign=\"bottom\"\n                                                    style=\"width:100%; max-width:100%; padding:0cm 0cm 0cm 0cm\">\n                                                    <div align=\"center\">\n                                                        <table class=\"x_$title_and_description_block\" border=\"0\"\n                                                               cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                                                               style=\"width:100%; max-width:100%; border-collapse:collapse; border-spacing:0px\">\n                                                            <tbody>\n                                                            <tr style=\"background-color:#444343; width:580px\">\n                                                                <td colspan=\"4\" style=\"padding:0px\">\n                                                                    <div width=\"100%\"\n                                                                         style=\"height:24px; line-height:21.8px; font-size:7px; width:100%\">\n                                                                        &nbsp;\n                                                                    </div>\n                                                                </td>\n                                                            </tr>\n                                                            </tbody>\n                                                        </table>\n                                                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                                                               style=\"width:100%; max-width:100%; border-collapse:collapse; border-spacing:0px\">\n                                                            <tbody>\n                                                            <tr style=\"background-color:#444343\">\n                                                                <td style=\"width:24px; max-width:24px; min-width:24px\">\n                                                                    &nbsp;\n                                                                </td>\n                                                                <td colspan=\"2\"\n                                                                    style=\"padding:0; color: white; font-style:inherit;  font-variant:inherit; font-weight:inherit; font-size:inherit; line-height:inherit\">\n                                                                    <h2 style=\"font-family:'MTS Text',sans-serif,serif,EmojiFont\">Привет!</h2>\n                                                                    <p style=\"font-family:'MTS Text',sans-serif,serif,EmojiFont\"></p>\n                                                                    <p style=\"font-family:'MTS Text',sans-serif,serif,EmojiFont\">\n                                                                        Рады сообщить, что мы успешно забронировали для Вас место в поезде согласно Вашей заявке на автобронирование!\n                                                                    </p>\n                                                                </td>\n                                                                <td style=\"width:24px; max-width:24px; min-width:24px\">\n                                                                    &nbsp;\n                                                                </td>\n                                                            </tr>\n                                                            </tbody>\n                                                        </table>\n                                                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                                                               style=\"width:100%; max-width:100%; border-collapse:collapse; border-spacing:0px\">\n                                                            <tbody>\n                                                            <tr style=\"background-color:#444343\">\n                                                                <td style=\"width:24px; max-width:24px; min-width:24px\">\n                                                                    &nbsp;\n                                                                </td>\n                                                                <td colspan=\"2\"\n                                                                    style=\"padding:0; font-style:inherit; font-variant:inherit; font-weight:inherit; font-size:inherit; line-height:inherit\">\n                                                                    <div width=\"532px\"\n                                                                         style=\"height:24px; line-height:21.8px; font-size:7px; width:532px\">\n                                                                        &nbsp;\n                                                                    </div>\n                                                                    <table align=\"center\">\n                                                                        <tbody>\n                                                                        <tr>\n                                                                            <td style=\"height: 38px; background: #e5ed58; border-radius: 6px;\">\n                                                                                <a href=\"https://i.postimg.cc/jSNLJGMF/Frame-432.png\"\n                                                                                   style=\"height: 38px; text-align: center; font-family: MTS Text, Arial, sans-serif, serif, EmojiFont;; font-size: 12px; line-height: 38px; text-decoration: none; padding: 0; display: block; border-radius: 4px;\">\n                                                                                    <span>&nbsp;&nbsp;</span>\n                                                                                    <span style=\"font-family: MTS Text, sans-serif; color: #434242; font-size: 12px; line-height: 38px; font-weight: bold; letter-spacing: 0.05em; -webkit-text-size-adjust:none;\">ПЕРЕЙТИ К ОПЛАТЕ</span>\n                                                                                    <span>&nbsp;&nbsp;</span>\n                                                                                </a>\n                                                                            </td>\n                                                                        </tr>\n                                                                        </tbody>\n                                                                    </table>\n                                                                </td>\n                                                                <td style=\"width:24px; max-width:24px; min-width:24px\">\n                                                                    &nbsp;\n                                                                </td>\n                                                            </tr>\n                                                            </tbody>\n                                                        </table>\n                                                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                                                               style=\"width:100%; max-width:100%; border-collapse:collapse; border-spacing:0px\">\n                                                            <tbody>\n                                                            <tr style=\"background-color:#444343; width:100%\">\n                                                                <td colspan=\"4\" style=\"padding:0px\">\n                                                                    <div width=\"580px\"\n                                                                         style=\"height:24px; line-height:21.8px; font-size:7px; width:580px\">\n                                                                        &nbsp;\n                                                                    </div>\n                                                                </td>\n                                                            </tr>\n                                                            </tbody>\n                                                        </table>\n                                                    </div>\n                                                </td>\n                                            </tr>\n                                            </tbody>\n                                        </table>\n                                    </div>\n                                </td>\n                                <td colspan=\"4\" style=\"width:24px; max-width:24px; min-width:24px\">&nbsp;</td>\n                            </tr>\n                            </tbody>\n                        </table>\n                        <table border=\"0\" cellspacing=\"0\" cellpadding=\"0\" width=\"100%\"\n                               style=\"width:100%; border-collapse:collapse\">\n                            <tbody>\n                            <tr style=\"background-color:\">\n                                <td width=\"580px\" style=\"height:24px; line-height:21.8px; width:580px\">&nbsp;</td>\n                            </tr>\n                            </tbody>\n                        </table>\n                    </div>\n                </div>\n\n            </div>\n        </div>\n    </div>\n    <div class=\"_rp_c5\" style=\"display: none;\"></div>\n</div>\n</body>\n</html>"

	// Формирование заголовков и тела письма
	headers := make(map[string]string)
	headers["From"] = senderEmail
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + htmlBody

	// Установка аутентификации
	auth := smtp.PlainAuth("", senderEmail, password, smtpHost)

	// Отправка письма
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		senderEmail,
		[]string{to},
		[]byte(message),
	)

	return err
}
