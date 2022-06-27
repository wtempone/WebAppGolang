<template>
  <div class="hello">

    <div id="map"></div>

    <div>
      <div class="text-center">
        <v-progress-linear class="mx-2 my-2" :value="animationProgress * 100"></v-progress-linear>
        <span class="h3">
          {{ velocityPlus }}
        </span>
        <v-icon v-if="reverseMode" large color="grey darken-2">
          mdi-chevron-left-box-outline
        </v-icon>
        <v-icon v-if="!reverseMode" large color="grey darken-2">
          mdi-chevron-right-box-outline
        </v-icon>
        <v-btn @click="backward()" class="mx-2" fab dark color="indigo">
          <v-icon dark>
            mdi-skip-backward
          </v-icon>
        </v-btn>
        <v-btn @click="replay()" class="mx-2" fab dark color="indigo">
          <v-icon dark>
            mdi-replay
          </v-icon>
        </v-btn>
        <v-btn @click="play()" class="mx-2" fab dark color="indigo">
          <v-icon dark>
            mdi-play
          </v-icon>
        </v-btn>
        <v-btn @click="pause()" class="mx-2" fab dark color="indigo">
          <v-icon dark>
            mdi-pause
          </v-icon>
        </v-btn>
        <v-btn @click="forward()" class="mx-2" fab dark color="indigo">
          <v-icon dark>
            mdi-skip-forward
          </v-icon>
        </v-btn>
      </div>
    </div>


  </div>
</template>

<script>
import ThreeJSOverlayView from '@ubilabs/threejs-overlay-view';
import { CatmullRomCurve3, Vector3 } from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';
import { Line2 } from 'three/examples/jsm/lines/Line2.js';
import { LineMaterial } from 'three/examples/jsm/lines/LineMaterial.js';
import { LineGeometry } from 'three/examples/jsm/lines/LineGeometry.js';
import { getMapsApiOptions, loadMapsApi } from '../jsm/load-maps-api';
import { Mesh, MeshStandardMaterial, Curve, ExtrudeGeometry, Shape, PerspectiveCamera } from 'three';
import json from '../assets/609705.json'
//import CAR_MODEL_URL from 'url: ../assets/lowpoly-sedan.glb';

export default {
  name: 'Car',
  data: () => ({
    playing: false,
    reverseMode: false,
    velocityPlus: 1,
    CAR_FRONT: new Vector3(0, 1, 0),
    animationProgress: 0,
    deltaTime: 0.0005,
    VIEW_PARAMS: {
      center: {},
      zoom: 17,
      heading: 40,
      tilt: 65,
    },
    ANIMATION_DURATION: 30000,
    ANIMATION_POINTS: [
      { lat: 53.554473, lng: 10.008226 },
      { lat: 53.554913, lng: 10.008124 },
      { lat: 53.554986, lng: 10.007928 },
      { lat: 53.554775, lng: 10.006363 },
      { lat: 53.554674, lng: 10.006383 },
      { lat: 53.554473, lng: 10.006681 },
      { lat: 53.554363, lng: 10.006971 },
      { lat: 53.554453, lng: 10.008091 },
      { lat: 53.554424, lng: 10.008201 },
      { lat: 53.554473, lng: 10.008226 }
    ],
    mapContainer: document.querySelector('#map'),
    tmpVec3: new Vector3(),
    jsonTrack: json,
  }),
  props: {
    msg: String
  },
  methods: {
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

    createScene: async function () {

      if (this.jsonTrack.trackingLog.coords[0]) {
        const cord = this.jsonTrack.trackingLog.coords[0]
        this.VIEW_PARAMS.center = { lat: cord.lat, lng: cord.lng }
      } else {
        console.log("Coordenadas nao iniciadas")
      }

      const map = await this.initMap();
      console.log('parameteres:', this.VIEW_PARAMS)
      const overlay = new ThreeJSOverlayView(this.VIEW_PARAMS.center);
      const scene = overlay.getScene();


      overlay.setMap(map);

      // create a Catmull-Rom spline from the points to smooth out the corners
      // for the animation

      var points
      if (this.jsonTrack.trackingLog.coords) {
        const cords = this.jsonTrack.trackingLog.coords
        const origin = cords[0].altitude
        const cordsAjusted = cords.map(x => ({ lat: x.lat, lng: x.lng, altitude: 0 }))
        points = cordsAjusted.map(p => overlay.latLngAltToVector3(p));
      }

      const curve = new CatmullRomCurve3(points, false, 'catmullrom', 0.2);
      curve.updateArcLengths();

      const trackLine = this.createTrackLine(curve);

      scene.add(trackLine);

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
        depth: 2,  // ui: depth
        bevelEnabled: true,  // ui: bevelEnabled
        bevelThickness: 1,  // ui: bevelThickness
        bevelSize: 1,  // ui: bevelSize
        bevelSegments: 2,  // ui: bevelSegments
      };

      const geometry = new ExtrudeGeometry(shape, extrudeSettings);

      const carModel = new Mesh(
        geometry,
        new MeshStandardMaterial({ color: 0xff0000 })
      );
      scene.add(carModel);
      overlay.requestRedraw();
      // the update-function will animate the car along the spline
      overlay.update = () => {
        // console.log("inicio update => ")
        trackLine.material.resolution.copy(overlay.getViewportSize());

        if (!carModel) return;
        if (this.playing) {
          if (this.animationProgress > 1) {
            this.animationProgress = 1
          } else if (this.animationProgress < 0) {
            this.animationProgress = 0
          } else {
            //  (performance.now() % this.ANIMATION_DURATION) / this.ANIMATION_DURATION;
            curve.getPointAt(this.animationProgress, carModel.position);
            curve.getTangentAt(this.animationProgress, this.tmpVec3);
            carModel.quaternion.setFromUnitVectors(this.CAR_FRONT, this.tmpVec3);

            const indexPosition = parseInt(this.jsonTrack.trackingLog.coords.length * this.animationProgress);
            const cameraPosition = this.jsonTrack.trackingLog.coords[indexPosition] 
            
            map.moveCamera({
              center: new google.maps.LatLng(cameraPosition.lat, cameraPosition.lng),
            });


          }
          this.animationProgress =
            this.animationProgress + ((this.reverseMode ? -this.deltaTime : this.deltaTime) * this.velocityPlus)
        }
        
        overlay.requestRedraw();
        // console.log(
        //   "final update => animationProgress:", this.animationProgress,
        //   ", playing:", this.playing,
        //   ", reverseMode:", this.reverseMode,
        //   ", deltaTime:", this.deltaTime,
        //   ", velocityPlus:", this.velocityPlus,
        // )

      };
    },

    /**
     * Load the Google Maps API and create the fullscreen map.
     */
    initMap: async function () {
      const { mapId } = getMapsApiOptions();
      await loadMapsApi();

      return new google.maps.Map(document.querySelector('#map'), {
        mapId,
        disableDefaultUI: true,
        backgroundColor: 'transparent',
        gestureHandling: 'greedy',
        mapTypeId: google.maps.MapTypeId.SATELLITE,
        ...this.VIEW_PARAMS
      });
    },

    /**
     * Create a mesh-line from the spline to render the track the car is driving.
     */
    createTrackLine: function (curve) {
      const numPoints = 10 * curve.points.length;
      const curvePoints = curve.getSpacedPoints(numPoints);
      const positions = new Float32Array(numPoints * 3);

      for (let i = 0; i < numPoints; i++) {
        curvePoints[i].toArray(positions, 3 * i);
      }

      const trackLine = new Line2(
        new LineGeometry(),
        new LineMaterial({ color: 0xFF0000 })
      );

      trackLine.geometry.setPositions(positions);

      return trackLine;
    },

    /**
     * Load and prepare the car-model for animation.
     */
    loadCarModel: async function () {
      const loader = new GLTFLoader();

      return new Promise(resolve => {
        loader.load(CAR_MODEL_URL, gltf => {
          const group = gltf.scene;
          const carModel = group.getObjectByName('sedan');

          carModel.scale.setScalar(3);
          carModel.rotation.set(Math.PI / 2, 0, Math.PI, 'ZXY');

          resolve(group);
        });
      });
    }
  },
  async mounted() {
    this.createScene()
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

#map {
  position: relative;
  top: 0;
  left: 0;
  width: 100%;
  height: 500px;
  z-index: 0;
}
</style>
