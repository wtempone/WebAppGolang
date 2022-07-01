import { PerspectiveCamera  } from 'three';
const createPerspectiveCamera = function () {
    const camera = new PerspectiveCamera(
        75,
        window.innerWidth / window.innerHeight,
        0.1,
        1000
    );
    camera.position.z = 500;
    return camera;
}

export {
    createPerspectiveCamera
}