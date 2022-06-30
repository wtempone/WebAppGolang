<template>

    <div>
        <Map3DScene :jsonTrack="jsonTrack" ref="map3DScene"
        ></Map3DScene>
<!--  -->
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

                <v-btn @click="$refs.map3DScene.forward()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-skip-forward
                    </v-icon>
                </v-btn>

                <v-btn @click="$refs.map3DScene.pause()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-pause
                    </v-icon>
                </v-btn>

                <v-btn @click="$refs.map3DScene.play()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-play
                    </v-icon>
                </v-btn>

                <v-btn @click="$refs.map3DScene.replay()" class="mx-2" small fab dark color="indigo">
                    <v-icon dark>
                        mdi-replay
                    </v-icon>
                </v-btn>

                <v-btn @click="$refs.map3DScene.backward()" class="mx-2" small fab dark color="indigo">
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

                <v-card v-if="$refs.map3DScene && $refs.map3DScene.jsonTrack.trackingLog"  class="mx-auto telemetry-card" max-width="400">

                    <v-list-item >
                        <v-list-item-icon>
                            <v-icon color="white">mdi-send</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{
                                $refs.map3DScene.jsonTrack.trackingLog.coords[$refs.map3DScene.animationProgressIndex].altitude
                        }}
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-cloud-download</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{ $refs.map3DScene.jsonTrack.trackingLog.times[$refs.map3DScene.animationProgressIndex]
                        }}
                        </v-list-item-subtitle>
                    </v-list-item>

                    <v-list-item>
                        <v-list-item-icon>
                            <v-icon color="white">mdi-cloud-download</v-icon>
                        </v-list-item-icon>
                        <v-list-item-subtitle class="white--text">{{
                                $refs.map3DScene.jsonTrack.trackingLog.climbs[$refs.map3DScene.animationProgressIndex]
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

import moment from 'moment';
import Map3DScene from './Map3DScene.vue';
import json from '../assets/609705.json'

export default {
    components: { Map3DScene },
    data: () => ({
        jsonTrack: json,
        viewData: {},
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
    methods: {

    },
    async mounted() {
    }
}
</script>
<style  lang= scss>
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