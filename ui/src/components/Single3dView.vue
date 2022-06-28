<template>
    <div>
        <div id="appgl" class="root3dElement">
        </div>

        
        <div class="progress-panel">

            <div class="labels">
                <v-icon v-if="reverseMode" small color="white">
                    mdi-chevron-left-box-outline
                </v-icon>

                <span color="white">
                    {{ velocityPlus }}
                </span>

                <v-icon v-if="!reverseMode" small color="white">
                    mdi-chevron-right-box-outline
                </v-icon>
            </div>
            <v-progress-linear class="mx-0"></v-progress-linear>

        </div>
    </div>
</template>
<script>

import {
    Mesh,
    PerspectiveCamera,
    WebGLRenderer,
    Scene,
    Color,
    BoxGeometry,
    MeshLambertMaterial,
    PointLight,
    Shape,
    ExtrudeGeometry,
    MeshStandardMaterial,
    Vector3,
    CatmullRomCurve3
} from 'three';

import { InteractionManager } from "three.interactive";
import * as TWEEN from "@tweenjs/tween.js";
import ThreeJSOverlayView from '@ubilabs/threejs-overlay-view';
import { Line2 } from 'three/examples/jsm/lines/Line2.js';
import { LineMaterial } from 'three/examples/jsm/lines/LineMaterial.js';
import { LineGeometry } from 'three/examples/jsm/lines/LineGeometry.js';
import {
    getMapsApiOptions,
    loadMapsApi
} from '../jsm/load-maps-api';
import moment from 'moment';

export default {
    props:{
        jsonTrack: {},
    },
    data: () => ({
        map: {},
        idRoot3dElement: "appgl",
        renderer: {},
        scene: {},
        camera: {},
        interactionManager: {},
        cubes: [],
        light: {},
        overlay: {},
        VIEW_PARAMS: {
            center: { },
            zoom: 17,
            heading: 40,
            tilt: 65,
        },
        playing: false,
        reverseMode: false,
        velocityPlus: 1,
        PLAYER_FRONT: new Vector3(0, 1, 0),
        tmpVec3: new Vector3(),
        animationProgress: 0,
        animationProgressIndex: 0,
        deltaTime: 0.0005,
        curve: {},
        trackLine: {},
        player: {},
        timeToUpdateCAmera: 10000,
        time: {},
    }),
    watch: {

    },
    methods: {
        velocityUp: function () {
            if (this.velocityPlus >= 6) re
            {
                this.velocityPlus++
            }
        },
        dateTime(value) {
            return moment(value).format('H:mm');
        },
        formatHour(value) {
            return moment(value, "HH:mm:ss").format('H:mm');
        },
        velocityDown: function () {
            if (this.velocityPlus < 2) return
            this.velocityPlus--
        },
        backward: function () {
            if (this.reverseMode) {
                this.velocityUp()
            } else if (this.velocityPlus > 1) {
                this.velocityDown()
            } else {
                this.reverseMode = true
            }
        },
        replay: function () {
            this.reverseMode = true
            this.velocityPlus = 1
            this.playing = true
        },
        play: function () {
            this.reverseMode = false
            this.velocityPlus = 1
            this.playing = true
        },
        pause: function () {
            this.playing = false
        },
        forward: function () {
            if (!this.reverseMode) {
                this.velocityUp()
            } else if (this.velocityPlus > 1) {
                this.velocityDown()
            } else {
                this.reverseMode = false
            }
        },
        animate: function (callback) {
            const modify = this.modifyScene
            function loop(time) {

                modify(time)

                callback(time);
                requestAnimationFrame(loop);
            }
            requestAnimationFrame(loop);
        },
        modifyScene: function (time) {
            this.trackLine.material.resolution.copy(this.overlay.getViewportSize());

            if (!this.player) return;
            if (this.playing) {
                if (this.animationProgress > 1) {
                    this.animationProgress = 1
                } else if (this.animationProgress < 0) {
                    this.animationProgress = 0
                } else {
                    //  (performance.now() % this.ANIMATION_DURATION) / this.ANIMATION_DURATION;
                    this.curve.getPointAt(this.animationProgress, this.player.position);
                    this.curve.getTangentAt(this.animationProgress, this.tmpVec3);
                    this.player.quaternion.setFromUnitVectors(this.PLAYER_FRONT, this.tmpVec3);
                    if (this.time < this.time * this.timeToUpdateCAmera) {
                        this.time = this.time + time
                    } else {
                        this.timeToUpdateCAmera = 0
                        const indexPosition = parseInt(this.jsonTrack.trackingLog.coords.length * this.animationProgress);
                        this.animationProgressIndex = indexPosition;
                        const cameraPosition = this.jsonTrack.trackingLog.coords[indexPosition]

                        const cameraOptions = {
                            tilt: this.VIEW_PARAMS.tilt,
                            heading: this.VIEW_PARAMS.heading,
                            zoom: this.VIEW_PARAMS.zoom,
                            center: this.VIEW_PARAMS.center,
                        };
                        new TWEEN.Tween(cameraOptions) // Create a new tween that modifies 'cameraOptions'.
                            .to(
                                {
                                    tilt: this.VIEW_PARAMS.tilt,
                                    heading: this.VIEW_PARAMS.heading,
                                    zoom: this.VIEW_PARAMS.zoom,
                                    center: { lat: cameraPosition.lat, lng: cameraPosition.lng },

                                },
                                time * this.timeToUpdateCAmera) // Move to destination in 15 second.
                            .easing(TWEEN.Easing.Quadratic.Out) // Use an easing function to make the animation smooth.
                            .onUpdate(() => {
                                this.map.moveCamera(cameraOptions);
                            })
                            .start(); // Start the tween immediately.
                    }

                }
                this.animationProgress =
                    this.animationProgress + ((this.reverseMode ? -this.deltaTime : this.deltaTime) * this.velocityPlus)
            }
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
        createHeartShape: function () {

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
                depth: 20,  // ui: depth
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
        },
        createLight: function () {
            const light = new PointLight(0xffffff, 1, 1000);
            light.position.set(0, 0, 10);
            return light;
        },
        createRenderer: function () {
            const root = document.getElementById(this.idRoot3dElement);
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
        config3dSpace: async function () {

            if (this.jsonTrack.trackingLog.coords[0]) {
                const cord = this.jsonTrack.trackingLog.coords[0]
                this.VIEW_PARAMS.center = { lat: cord.lat, lng: cord.lng }
                this.jsonTrack.trackingLog.times = this.jsonTrack.trackingLog.times.map(x => (this.dateTime(x)))
                this.jsonTrack.voo.horaDecolagem = this.formatHour(this.jsonTrack.voo.horaDecolagem)
                this.jsonTrack.voo.horaPouso = this.formatHour(this.jsonTrack.voo.horaPouso)
                this.jsonTrack.voo.duracao = this.formatHour(this.jsonTrack.voo.duracao)
            } else {
                console.log("Coordenadas nao iniciadas")
            }

            this.renderer = this.createRenderer();
            this.scene = this.createScene();
            this.camera = this.createCamera();

            this.interactionManager = new InteractionManager(
                this.renderer,
                this.camera,
                this.renderer.domElement
            );

            this.map = await this.initMap()
            this.addMap(this.map)
            this.curve = this.convertTrackinglogToCurve()

            this.trackLine = this.createTrackLine(this.curve);

            this.scene.add(this.trackLine);
            this.player = this.createHeartShape()
            this.scene.add(this.player)
            this.light = this.createLight()
            this.scene.add(this.light);

            this.overlay.requestRedraw();
            this.renderer.render(this.scene, this.camera);
            this.renderer.resetState();

            this.animate((time) => {
                this.overlay.requestRedraw();
                this.renderer.render(this.scene, this.camera);
                this.renderer.resetState();
                this.interactionManager.update();
                TWEEN.update(time);
            });
        },
        createTrackLine: function (curve) {
            const numPoints = 10 * curve.points.length;
            const curvePoints = curve.getSpacedPoints(numPoints);
            const positions = new Float32Array(numPoints * 3);

            for (let i = 0; i < numPoints; i++) {
                curvePoints[i].toArray(positions, 3 * i);
            }

            const trackLine = new Line2(
                new LineGeometry(),
                new LineMaterial({
                    color: 0xFF0000, linewidth: 0.01
                })
            );

            trackLine.geometry.setPositions(positions);

            return trackLine;
        },
        convertTrackinglogToCurve: function () {
            var points
            if (this.jsonTrack.trackingLog.coords) {
                const cords = this.jsonTrack.trackingLog.coords
                const origin = cords[0].altitude
                const cordsAjusted = cords.map(x => ({ lat: x.lat, lng: x.lng, altitude: 0 }))
                points = cordsAjusted.map(p => this.overlay.latLngAltToVector3(p));
            }

            const curve = new CatmullRomCurve3(points, false, 'catmullrom', 0.2);
            curve.updateArcLengths();
            return curve
        },
        initMap: async function () {
            const { mapId } = getMapsApiOptions();

            await loadMapsApi();

            return new google.maps.Map(document.getElementById(this.idRoot3dElement), {
                mapId,
                disableDefaultUI: true,
                backgroundColor: 'transparent',
                gestureHandling: 'greedy',
                mapTypeId: google.maps.MapTypeId.SATELLITE,
                ...this.VIEW_PARAMS
            });
        },
        addMap: function (map) {
            this.overlay = new ThreeJSOverlayView(this.VIEW_PARAMS.center);
            this.scene = this.overlay.getScene();
            this.overlay.setMap(map);
        }
    },
    async mounted() {
        this.config3dSpace()
    }
}
</script>
<style  lang= scss>
.root3dElement {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 0;
}
</style>