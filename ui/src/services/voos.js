import httpClient from '@/services/httpClient';

const END_POINT = '/voo';

const obterListaVoos = () => httpClient.get(END_POINT);
const obterVoo = (id) => httpClient.get(`${END_POINT}/${id}`);

export {
    obterListaVoos,
    obterVoo
}