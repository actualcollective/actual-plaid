import './style.css';
import Alpine from 'alpinejs';
import connect from './connect';
import plaid from "./plaid";

window.Alpine = Alpine;

Alpine.store('error', '');
Alpine.store('isLoading', true);

// @ts-ignore
Alpine.data('connect', connect);
// @ts-ignore
Alpine.data('plaid', plaid)

Alpine.start();
