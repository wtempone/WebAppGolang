import { ExtrudeGeometry, MeshStandardMaterial, Shape, Mesh } from 'three';
const createHeartShape = function () {

    const shape = new Shape();
    const x = -2.5;
    const y = -5;
    shape.moveTo(x + 2.5, y + 2.5);
    shape.bezierCurveTo(x + 2.5, y + 2.5, x + 2, y, x, y);
    shape.bezierCurveTo(x - 3, y, x - 3, y + 3.5, x - 3, y + 3.5);
    shape.bezierCurveTo(x - 3, y + 5.5, x - 1.5, y + 7.7, x + 2.5, y + 9.5);
    shape.bezierCurveTo(x + 6, y + 7.7, x + 8, y + 4.5, x + 8, y + 3.5);
    shape.bezierCurveTo(x + 8, y + 3.5, x + 8, y, x + 5, y);
    shape.bezierCurveTo(x + 3.5, y, x + 2.5, y + 2.5, x + 2.5, y + 2.5);

    const extrudeSettings = {
        steps: 2,  // ui: steps
        depth: 4,  // ui: depth
        bevelEnabled: true,  // ui: bevelEnabled
        bevelThickness: 1,  // ui: bevelThickness
        bevelSize: 1,  // ui: bevelSize
        bevelSegments: 2,  // ui: bevelSegments
    };

    const geometry = new ExtrudeGeometry(shape, extrudeSettings);

    const mesh = new Mesh(
        geometry,
        new MeshStandardMaterial({ color: 0xff0000 })
    );
    return mesh;
}
export {
    createHeartShape  
}