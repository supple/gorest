package events_procesor

type EventProcessor struct {

}

type Event struct {

}

type Source interface {
    Pop() Event
    Ack(msgId string)
}

func Process() {

}
