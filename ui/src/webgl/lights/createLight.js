import { PointLight } from 'three';
const createligth = function () {
    const light = new PointLight(0xffffff, 1, 1000);
    light.position.set(0, 0, 10);
    return light;
}
export{
    createligth
}