import { fromEvent } from 'rxjs';

import { MESSAGE } from './const';

fromEvent(document, 'click').subscribe(() => console.log('message: %s', MESSAGE));
