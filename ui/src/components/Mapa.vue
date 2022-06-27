<template>
    <v-container fluid>
        <v-row>
            <v-col cols="12">
                <v-card>
                    <v-subheader>Detalhes</v-subheader>
                    <v-card-text>

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

        }
    },
    props: {
    },
    methods: {

    },
    mounted() {
        console.log('Pagina')

    }
}

</script>