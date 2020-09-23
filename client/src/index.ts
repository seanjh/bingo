import { fromEvent } from 'rxjs';

fromEvent(document, 'click').subscribe(() => console.log('click!'));
