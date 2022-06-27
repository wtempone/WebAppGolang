<template>
    <v-container fluid>
        <v-row>
            <v-col cols="12">
                <v-card>
                    <v-subheader>Detalhes</v-subheader>
                    <v-card-text>
                        <gmap-map :zoom="14" :center="center" style="width:100%;  height: 600px;">
                            <gmap-marker :key="index" v-for="(m, index) in locationMarkers" :position="m.position"
                                @click="center = m.position"></gmap-marker>
                        </gmap-map>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>

</template>
<script>
import { obterVoo } from "@/services/voos"
import { consoleError } from "vuetify/lib/util/console"

export default {
    data() {
        return {
            voo: {},
            center: {
                lat: 39.7837304,
                lng: -100.4458825
            },
            locationMarkers: [],
            locPlaces: [],
            existingPlace: null
        }
    },
    props: {
        link: {}
    },
    methods: {
        obter(link) {
            obterVoo(link).then(res => {
                this.voo = res.data
                console.log('detelhes', res.data)

            }).catch((err) => {
                console.log("err", err)
            })
        },
        initMarker(loc) {
            this.existingPlace = loc;
        },
        addLocationMarker() {
            if (this.existingPlace) {
                const marker = {
                    lat: this.existingPlace.geometry.location.lat(),
                    lng: this.existingPlace.geometry.location.lng()
                };
                this.locationMarkers.push({ position: marker });
                this.locPlaces.push(this.existingPlace);
                this.center = marker;
                this.existingPlace = null;
            }
        },
        locateGeoLocation: function () {
            navigator.geolocation.getCurrentPosition(res => {
                this.center = {
                    lat: res.coords.latitude,
                    lng: res.coords.longitude
                };
            });
        }
    },
    mounted() {
        console.log('passa', this.link)
        this.obter(this.link)
        this.locateGeoLocation();

    }
}

</script>