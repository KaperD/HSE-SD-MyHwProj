#!/usr/bin/env python3

import pika
import sys
import os
import json

import runner


def main():
    def callback(ch, method, properties, body):
        print(f" [x] Received: {body}")
        message = json.loads(body)
        homework = runner.Homework(message['Homework']['check'])
        submission = runner.Submission(message['Submission']['solution'])
        feedback = homework.get_feedback(submission)
        response = message['Submission']
        response['mark'] = feedback.mark
        response['comment'] = feedback.comment
        print(f" [x] Response build: {response}")
        ch.basic_publish(exchange='',
                         routing_key=properties.reply_to,
                         properties=pika.BasicProperties(
                             correlation_id=properties.correlation_id
                         ),
                         body=str(json.dumps(response)))
        ch.basic_ack(delivery_tag=method.delivery_tag)

    connection = pika.BlockingConnection(pika.ConnectionParameters(
        host='localhost',
        port=5672)
    )
    channel = connection.channel()
    channel.queue_declare(queue='rpc_queue')
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue='rpc_queue', on_message_callback=callback)

    print(" [x] Awaiting RPC requests")
    channel.start_consuming()


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
