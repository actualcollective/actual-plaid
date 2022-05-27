import {Plaid} from "./plaid-link";

interface Config {
  url: string;
  plaid: PlaidConfig;
}

interface PlaidConfig {
  clientName: string;
  product: Plaid.Product;
  env: Plaid.Environment;
  lang: Plaid.Language;
  countries: Plaid.Country[];
}

export default <Config>{
  url: 'http://localhost:8081',
  plaid: {
    clientName: 'Actual',
    product: 'transactions' ,
    env: 'sandbox',
    lang: 'en',
    countries: ['US']
  }
};
