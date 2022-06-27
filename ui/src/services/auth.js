import httpClient from '@/services/httpClient';

const END_POINT = '/auth';

const login = (email, senha) => httpClient.post(`${END_POINT}/login`, { email, senha });
const registro = (nome, email, senha) => httpClient.post(`${END_POINT}/registro`, {nome, email, senha });
const logout = () => httpClient.post(`${END_POINT}/logout`);
const check = () => httpClient.get(`${END_POINT}/check`);

export {
    login,
    registro,
    logout,
    check
}