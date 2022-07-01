import { OrthographicCamera } from 'three';

const createOrthographicCamera = function() {
    const camera =  new OrthographicCamera( window.width / - 2, window.width / 2, window.height / 2, window.height / - 2, 1, 1000 );
    camera.position.z = 500;
    return camera;
}
export {
    createOrthographicCamera
}