import httpClient from '@/services/httpClient';

const END_POINT = '/usuario';

const obterUsuario = () => httpClient.get(END_POINT);

export {
    obterUsuario
}