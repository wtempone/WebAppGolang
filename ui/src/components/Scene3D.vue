<template>
    <div id="appgl">

    </div>
</template>
<script>
import { Mesh, PerspectiveCamera, WebGLRenderer ,Scene , Color ,BoxGeometry, MeshLambertMaterial ,PointLight  } from 'three';
import { InteractionManager } from "three.interactive";
import * as TWEEN from "@tweenjs/tween.js";

export default {
    data: () => ({
        renderer: {},
        scene: {},
        camera: {},
        interactionManager: {},
        cubes: [],
        light: {}
    }),
    methods: {
        animate: function (callback) {
            function loop(time) {
                callback(time);
                requestAnimationFrame(loop);
            }
            requestAnimationFrame(loop);
        },
        createCamera: function () {
            const camera = new PerspectiveCamera(
                75,
                window.innerWidth / window.innerHeight,
                0.1,
                1000
            );
            camera.position.z = 5;
            return camera;
        },
        createCube: function ({ color, x, y }) {
            const geometry = new BoxGeometry();
            const material = new MeshLambertMaterial({ color });
            const cube = new Mesh(geometry, material);
            cube.position.set(x, y, 0);

            return cube;
        },
        createLight: function () {
            const light = new PointLight(0xffffff, 1, 1000);
            light.position.set(0, 0, 10);
            return light;
        },
        createRenderer: function () {
            const root = document.getElementById("appgl");
            const renderer = new WebGLRenderer({ antialias: true });
            renderer.setSize(window.innerWidth, window.innerHeight);
            root.appendChild(renderer.domElement);
            return renderer;
        },
        createScene: function () {
            const scene = new Scene();
            scene.background = new Color(0xffffff);
            return scene;
        },
        config: function () {
            this.renderer = this.createRenderer();
            this.scene = this.createScene();
            this.camera = this.createCamera();
            this.interactionManager = new InteractionManager(
                this.renderer,
                this.camera,
                this.renderer.domElement
            );

            this.cubes = {
                pink: this.createCube({ color: 0xff00ce, x: -1, y: -1 }),
                purple: this.createCube({ color: 0x9300fb, x: 1, y: -1 }),
                blue: this.createCube({ color: 0x0065d9, x: 1, y: 1 }),
                cyan: this.createCube({ color: 0x00d7d0, x: -1, y: 1 })
            };

            this.light = this.createLight();

            for (const [name, object] of Object.entries(this.cubes)) {
                object.addEventListener("click", (event) => {
                    event.stopPropagation();
                    console.log(`${name} cube was clicked`);
                    const cube = event.target;
                    const coords = { x: this.camera.position.x, y: this.camera.position.y };
                    new TWEEN.Tween(coords)
                        .to({ x: cube.position.x, y: cube.position.y })
                        .easing(TWEEN.Easing.Quadratic.Out)
                        .onUpdate(() =>
                            this.camera.position.set(coords.x, coords.y, this.camera.position.z)
                        )
                        .start();
                });
                this.interactionManager.add(object);
                this.scene.add(object);
            }

            this.scene.add(this.light);

            this.animate((time) => {
                this.renderer.render(this.scene, this.camera);
                this.interactionManager.update();
                TWEEN.update(time);
            });
        }
    },

    async mounted() {
        this.config()
    }
}
</script>
<style>
</style>