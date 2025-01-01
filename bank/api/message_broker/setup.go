package message_broker

import (
	"log"

	"github.com/onlineTraveling/bank/api/message_broker/handler"
	"github.com/onlineTraveling/bank/app"
)

func Run(app *app.App) {
	messageBroker := app.MessageBroker()
	bankHandler := handler.NewBankHandler(app.BankService())
	createWalletQueueName := app.GetConfig().MessageBroker.CreateWalletQueueName
	transferTransactionQueueName := app.GetConfig().MessageBroker.TransferQueueName
	log.Println("Broker is started..")
	go messageBroker.Consume(createWalletQueueName, bankHandler.CreateWallet)
	go messageBroker.Consume(transferTransactionQueueName, bankHandler.Transfer)
}
