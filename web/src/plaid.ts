import { Plaid } from './plaid-link';
import axios from "axios";
import config from "./config";

export default () => ({
    linked: false,
    link(token: string, bankCtx: string, plaidToken: string) {
        const product: Plaid.Product = config.plaid.product;
        const env: Plaid.Environment = config.plaid.env;
        const language: Plaid.Language = config.plaid.lang;
        const countryCodes: Plaid.Country[] = config.plaid.countries;

        const plaidConfig: Plaid.CreateConfig = {
            clientName: config.plaid.clientName,
            product: [product],
            token: plaidToken,
            env,
            language,
            countryCodes,
            onSuccess: async (publicToken, metadata) => {
                this.linked = true;
                const res = await axios.post(config.url + '/api/plaid/success', { token: token, bankCtx: bankCtx, publicToken: publicToken, metadata: metadata });
                console.log(res.data)
            }
        }

        const handler = window.Plaid.create(plaidConfig);
        handler.open();
    },
})