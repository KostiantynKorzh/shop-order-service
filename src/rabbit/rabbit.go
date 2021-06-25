package rabbit

const RABBIT_URL = "amqp://guest:guest@localhost:5672/"

//func Init() *amqp.Channel {
//
//	rabbitCon, err := amqp.Dial(RABBIT_URL)
//	if err != nil {
//		panic(err)
//	}
//
//	rabbitChannel, err := rabbitCon.Channel()
//	if err != nil {
//		panic(err)
//	}
//
//	_, err = rabbitChannel.QueueDeclare(
//		"test-queue",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	return rabbitChannel
//}
