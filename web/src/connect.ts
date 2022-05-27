import axios from 'axios';
import config from './config';
import Alpine from 'alpinejs';

export default (options: string) => ({
  token: null,
  plaidToken: null,
  bankCtx: null,
  init() {
    Alpine.store('isLoading', true);
    this.send()
      .then((data) => {
        console.log(data)
        this.plaidToken = data.plaidToken;
        this.token = data.token;
        this.bankCtx = data.bankCtx;
        Alpine.store('isLoading', false);
      })
      .catch((err) => {
        console.log(err)
        Alpine.store('error', this.reason(err.response?.data?.reason));
        Alpine.store('isLoading', false);
      });
  },
  async send() {
    const res = await axios.post(config.url + '/api/install', { options: options.replaceAll(' ', '+') });
    return res.data;
  },
  reason(error: string) {
    let msg = 'unknown error occurred';
    switch (error) {
      default:
        break;
    }

    return msg;
  },
});
