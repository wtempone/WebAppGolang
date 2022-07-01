import { MeshLambertMaterial, BoxGeometry } from 'three';

const createCube = function ({ color, x, y }) {
    const geometry = new BoxGeometry();
    const material = new MeshLambertMaterial({ color });
    const cube = new Mesh(geometry, material);
    cube.position.set(x, y, 0);
    return cube;
}

export {
    createCube
}