import { fromEvent } from 'rxjs';

const eventSource = new EventSource('//localhost:8000/api/v1/subscribe/bar');
const eventSourceMessages = fromEvent(eventSource, 'onmessage');

eventSourceMessages.subscribe((event) => {
  console.log('event: %j', event);
  const elem = document.getElementById('output');
  if (elem) {
    elem.innerText = `${elem.innerText ?? ''}\n${event.toString()}`;
  }
});
