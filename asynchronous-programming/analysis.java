package com.lightbend.akka.sample;

import akka.actor.*;

class Ts {
    public String tsName;

    public Ts(String tsName) {
        this.tsName = tsName;
    }
}

class Nlp extends AbstractActor {

    static Props props() {
        return Props.create(Nlp.class, () -> new Nlp());
    }

    @Override
    public Receive createReceive() {
        return receiveBuilder()
                .match(Asr.Result.class, asr -> {
                    System.out.printf("nlp: %s\n", asr.text);
                    System.out.println("nlp processing... done");
                })
                .build();
    }
}

class Asr extends AbstractActor {
    private final ActorRef nlp;

    static Props props() {
        return Props.create(Asr.class, () -> new Asr());
    }

    public Asr() {
        this.nlp = context().actorOf(Nlp.props());
    }

    @Override
    public Receive createReceive() {
        return receiveBuilder()
                .match(Ts.class, ts -> {
                    System.out.println("asr processing... done");
                    this.nlp.tell(new Result("asr result"), getSelf());
                })
                .build();
    }

    class Result {
        String text;

        Result (String asr) {
            this.text= asr;
        }
    }
}

class Ocr extends AbstractActor {

    static Props props() {
        return Props.create(Ocr.class, () -> new Ocr());
    }

    @Override
    public Receive createReceive() {
        return receiveBuilder()
                .match(Ts.class, ts -> {
                    System.out.println("ocr processing... done");
                })
                .build();
    }
}

public class Analysis {
    public static void main(String[] args) {
        final ActorSystem system = ActorSystem.create("Analysis");

        ActorRef asrActor = system.actorOf(Asr.props(), "asr");
        ActorRef ocrActor = system.actorOf(Ocr.props(), "ocr");
        Ts ts = new Ts("this is a sample ts");
        asrActor.tell(ts, ActorRef.noSender());
        ocrActor.tell(ts, ActorRef.noSender());

        system.terminate();
    }
}

