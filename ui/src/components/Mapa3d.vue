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
        <div class="panel-controls">

            <v-speed-dial v-model="controlsSlide.fab" :top="controlsSlide.top" :bottom="controlsSlide.bottom"
                :right="controlsSlide.right" :left="controlsSlide.left" :direction="controlsSlide.direction"
                :open-on-hover="controlsSlide.hover" :transition="controlsSlide.transition" elevation="20">
                <template v-slot:activator>
                    <v-btn v-model="controlsSlide.fab" color="blue darken-2" dark fab>
                        <v-icon v-if="controlsSlide.fab">
                            mdi-close
                        </v-icon>
                        <v-icon v-else>
                            mdi-google-controller </v-icon>
                    </v-btn>
                </template>



                <v-btn @click="forward()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-skip-forward
                    </v-icon>
                </v-btn>

                <v-btn @click="pause()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-pause
                    </v-icon>
                </v-btn>

                <v-btn @click="play()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-play
                    </v-icon>
                </v-btn>

                <v-btn @click="replay()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-replay
                    </v-icon>
                </v-btn>

                <v-btn @click="backward()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-skip-backward
                    </v-icon>
                </v-btn>


            </v-speed-dial>




        </div>

        <div class="panel-telemetry">

            <v-speed-dial v-model="controlsTelemetry.fab" :top="controlsTelemetry.top"
                :bottom="controlsTelemetry.bottom" :right="controlsTelemetry.right" :left="controlsTelemetry.left"
                :direction="controlsTelemetry.direction" :open-on-hover="controlsTelemetry.hover"
                :transition="controlsTelemetry.transition" elevation="20">
                <template v-slot:activator>
                    <v-btn v-model="controlsTelemetry.fab" color="blue darken-2" dark fab>
                        <v-icon v-if="controlsTelemetry.fab">
                            mdi-close </v-icon>
                        <v-icon v-else>
                            mdi-gauge </v-icon>
                    </v-btn>
                </template>



                <v-card class="mx-auto telemetry-card" max-width="400">
                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-send</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{
                                jsonTrack.trackingLog.coords[animationProgressIndex].altitude
                        }}
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-cloud-download</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{ jsonTrack.trackingLog.times[animationProgressIndex]
                        }}
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-cloud-download</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{
                                jsonTrack.trackingLog.climbs[animationProgressIndex]
                        }}
                        </v-list-item-subtitle>
                    </v-list-item>
                </v-card>

            </v-speed-dial>

        </div>

        <div class="panel-info">

            <v-speed-dial v-model="controlsInfo.fab" :top="controlsInfo.top" :bottom="controlsInfo.bottom"
                :right="controlsInfo.right" :left="controlsInfo.left" :direction="controlsInfo.direction"
                :open-on-hover="controlsInfo.hover" :transition="controlsInfo.transition" elevation="20">
                <template v-slot:activator>
                    <v-btn v-model="controlsInfo.fab" color="blue darken-2" dark fab>
                        <v-icon v-if="controlsInfo.fab">
                            mdi-close </v-icon>
                        <v-icon v-else>
                            mdi-information-variant
                        </v-icon>
                    </v-btn>
                </template>



                <v-card class="mx-auto info-card" max-width="400">
                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-account-circle</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{ jsonTrack.voo.piloto }}
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-calendar</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{ jsonTrack.voo.data }}
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-airplane-takeoff</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{ jsonTrack.voo.decolagem }} ({{
                                jsonTrack.voo.horaDecolagem
                        }})
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-airplane-landing</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{ jsonTrack.voo.pouso }} ({{ jsonTrack.voo.horaPouso
                        }})
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-clock-end</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{ jsonTrack.voo.duracao }}
                        </v-list-item-subtitle>
                    </v-list-item>

                </v-card>

            </v-speed-dial>

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
import json from '../assets/496171.json'
import { Line2 } from 'three/examples/jsm/lines/Line2.js';
import { LineMaterial } from 'three/examples/jsm/lines/LineMaterial.js';
import { LineGeometry } from 'three/examples/jsm/lines/LineGeometry.js';
import {
    getMapsApiOptions,
    loadMapsApi
} from '../jsm/load-maps-api';
import moment from 'moment';

export default {
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
            center: { lat: 53.554486, lng: 10.007479 },
            zoom: 17,
            heading: 40,
            tilt: 65,
        },
        playing: false,
        reverseMode: false,
        velocityPlus: 1,
        PLAYER_FRONT: new Vector3(0, 1, 0),
        tmpVec3: new Vector3(),
        jsonTrack: json,
        animationProgress: 0,
        animationProgressIndex: 0,
        deltaTime: 0.0005,
        curve: {},
        trackLine: {},
        player: {},
        timeToUpdateCAmera: 10000,
        time: {},
        infoComponent: {},
        controlsSlide: {
            direction: 'left',
            fab: true,
            fling: false,
            hover: false,
            tabs: null,
            top: false,
            right: true,
            bottom: true,
            left: false,
            transition: 'slide-x-reverse-transition',
        },
        controlsTelemetry: {
            direction: 'bottom',
            fab: true,
            fling: false,
            hover: false,
            tabs: null,
            top: true,
            right: true,
            bottom: false,
            left: false,
            transition: 'slide-y-reverse-transition',
        },
        controlsInfo: {
            direction: 'bottom',
            fab: false,
            fling: false,
            hover: false,
            tabs: null,
            top: true,
            right: false,
            bottom: false,
            left: true,
            transition: 'slide-y-reverse-transition',
        },
        telemetryInfo: {
            labels: ['SU', 'MO', 'TU', 'WED', 'TH', 'FR', 'SA'],
            time: 0,
            forecast: [
                { day: 'Tuesday', icon: 'mdi-white-balance-sunny', temp: '24\xB0/12\xB0' },
                { day: 'Wednesday', icon: 'mdi-white-balance-sunny', temp: '22\xB0/14\xB0' },
                { day: 'Thursday', icon: 'mdi-cloud', temp: '25\xB0/15\xB0' },
            ],
        }
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
            this.controlsSlide.fab = false
            if (this.reverseMode) {
                this.velocityUp()
            } else if (this.velocityPlus > 1) {
                this.velocityDown()
            } else {
                this.reverseMode = true
            }
        },
        replay: function () {
            this.controlsSlide.fab = false
            this.reverseMode = true
            this.velocityPlus = 1
            this.playing = true
        },
        play: function () {
            this.controlsSlide.fab = false
            this.reverseMode = false
            this.velocityPlus = 1
            this.playing = true
        },
        pause: function () {
            this.controlsSlide.fab = false
            this.playing = false
        },
        forward: function () {
            this.controlsSlide.fab = false
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
            // console.log(
            //     "final update => animationProgress:", this.animationProgress,
            //     ", playing:", this.playing,
            //     ", reverseMode:", this.reverseMode,
            //     ", deltaTime:", this.deltaTime,
            //     ", velocityPlus:", this.velocityPlus,
            // )
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

.panel-controls {
    position: absolute;
    bottom: 0;
    right: 0;
    display: inline-block;
    padding-bottom: 40px;
    padding-right: 0px;
}

.panel-telemetry {
    position: absolute;
    top: 0;
    right: 0;
    display: inline-block;
    padding-top: 60px;
    padding-right: 0px;
}

.panel-info {
    position: absolute;
    top: 0;
    left: 0;
    display: inline-block;
    padding-top: 60px;
    padding-right: 0px;
}

.telemetry-card {

    left: -94px;
    width: 150px;
    background-color: rgba(50, 50, 50, 0.5) !important;
    color: white !important;
}

.info-card {

    left: 0;
    width: 42vh;
    margin-right: 40px;
    background-color: rgba(50, 50, 50, 0.5) !important;
    color: white !important;
}

.progress-panel {
    display: block;
    align-content: center;
    position: absolute;
    z-index: 100;
    width: 100%;
    bottom: 30px;
}

.labels {
    color: white;
    display: flex;
    justify-content: center;
    font-size: 15px;
}
</style>