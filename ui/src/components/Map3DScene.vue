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
    WebGLRenderer,
    Scene,
    Color,
    Vector3,
    CatmullRomCurve3,
    TextureLoader,
    PlaneBufferGeometry,
    MeshStandardMaterial,
    Mesh 
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
import { createHeartShape } from '@/webgl/objects.js'
import { createPerspectiveCamera, createOrthographicCamera } from '@/webgl/cameras.js'
import { createligth } from '@/webgl/lights.js'
import { getImageMap } from "@/services/google"
import moment from 'moment';
export default {
    props: {
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
        zoom: 17,
        VIEW_PARAMS: {
            center: { lat: -19.35065, lng: -42.56858 },
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
        cordsAjusted: {}
    }),

    watch: {

    },
    methods: {
        dateTime(value) {
            return moment(value).format('H:mm');
        },
        formatHour(value) {
            return moment(value, "HH:mm:ss").format('H:mm');
        },
        velocityUp: function () {
            if (this.velocityPlus >= 6) re
            {
                this.velocityPlus++
            }
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
            if (!this.player) return;
            if (this.playing) {
                this.trackLine.material.resolution.copy(this.overlay.getViewportSize());

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
        openFile: function () {
            if (this.jsonTrack.trackingLog.coords[0]) {
                const cord = this.jsonTrack.trackingLog.coords[0]
                this.VIEW_PARAMS.center = { lat: cord.lat, lng: cord.lng }

                this.jsonTrack.trackingLog.times = this.jsonTrack.trackingLog.times.map(x => (this.dateTime(x)))
                this.jsonTrack.voo.horaDecolagem = this.formatHour(this.jsonTrack.voo.horaDecolagem)
                this.jsonTrack.voo.horaPouso = this.formatHour(this.jsonTrack.voo.horaPouso)
                this.jsonTrack.voo.duracao = this.formatHour(this.jsonTrack.voo.duracao)
                this.animationProgress = 0
                this.time = 0
                this.playing = false
                this.curve = this.convertTrackinglogToCurve()
                this.trackLine = this.createTrackLine(this.curve);
                //  this.player.setPositions(this.cordsAjusted)
                this.scene.add(this.trackLine);
            }
        },
        createGround: async function () {
            const { apiKey, mapId } = getMapsApiOptions();
            const loader = new TextureLoader()
            const utlImage = getImageMap(this.VIEW_PARAMS.center, this.VIEW_PARAMS.zoom, 300, 2)
            const texture = loader.load(utlImage)

            const geometry = new PlaneBufferGeometry(300, 300, 20, 20)

            const positions = geometry.attributes.position.array;
            const material = new MeshStandardMaterial({ color: 0x999999, map: texture })

            var vertices = []
            for (let i = 0; i < positions.length; i += 3) {
                const v = new Vector3(positions[i], positions[i + 1], positions[i + 2])
                vertices.push(v)
            }

            let coords = vertices;
            coords = coords.map(p => (this.overlay.vector3ToLatLngAlt(p)));
            const coorsLatLng = coords.map(c => ({ lat: c.lat, lng: c.lng }))

            const altitudes = await this.getElevations(coorsLatLng);

            console.log('Diferenca = ', altitudes[0].elevation, this.jsonTrack.trackingLog.coords[0].altitude)

            var zpos = []
            for (let i = 0; i < coords.length; i++) {
                zpos.push(altitudes[i].elevation - this.jsonTrack.trackingLog.coords[0].altitude)
            }
            var zindex = 0
            for (let i = 0; i < positions.length; i += 3) {
                positions[i + 2] = zpos[zindex]
                zindex++
            }
            // console.log('positionsGeometry == >', positions)
            // geometry.positions = positions;

            const plane = new Mesh(geometry, material)

            return plane
        },
        updateCurrentLocation: function () {
            navigator.geolocation.getCurrentPosition(
                position => {
                    const coord = { lat: position.coords.latitude, lng: position.coords.longitude }
                    this.VIEW_PARAMS.center = { lat: coord.lat, lng: coord.lng }
                },
                error => {
                    console.log('erro - possition', error.message);
                },
            )
        },
        getElevations: async function (locations) {
            const elevator = new google.maps.ElevationService();
            return elevator
                .getElevationForLocations({
                    locations: locations,
                })
                .then(({ results }) => {
                    return results
                })
                .catch((e) =>
                    console.log("Elevation service failed due to: " + e)
                );

        },
        config3dSpace: async function () {
            //this.updateCurrentLocation();

            this.renderer = this.createRenderer();
            this.scene = this.createScene();
            // this.camera = createPerspectiveCamera();
            this.camera = createOrthographicCamera()

            this.interactionManager = new InteractionManager(
                this.renderer,
                this.camera,
                this.renderer.domElement
            );

            this.map = await this.initMap()
            this.addMap(this.map)

            const ground = await this.createGround()
            this.scene.add(ground)
            this.player = createHeartShape()
            this.scene.add(this.player)
            this.light = createligth()
            this.scene.add(this.light);

            this.openFile()


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
                    color: 0xFF0000, linewidth: 0.0001
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
                this.cordsAjusted = cords.map(x => ({ lat: x.lat, lng: x.lng, altitude: x.altitude - origin }))
                points = this.cordsAjusted.map(p => this.overlay.latLngAltToVector3(p));
            }

            const curve = new CatmullRomCurve3(points, false, 'catmullrom', 0.2);
            curve.updateArcLengths();
            return curve
        },
        initMap: async function () {
            const { apiKey, mapId } = getMapsApiOptions();

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