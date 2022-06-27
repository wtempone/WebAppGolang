<template>
    <v-container fluid>
        <v-row>
            <v-col cols="12">
                <v-card>
                    <v-subheader>Lista</v-subheader>

                    <v-list two-line v-if="lista">
                        <template v-for="item in lista">
                            <v-list-item @click="verVoo(item.link)">
                                <v-list-item-avatar color="grey darken-1">
                                </v-list-item-avatar>

                                <v-list-item-content>
                                    <v-list-item-title>{{ item.piloto }}</v-list-item-title>
                                    <v-list-item-subtitle>
                                        {{ item.decolagem }} - {{ item.data }}
                                    </v-list-item-subtitle>
                                    <v-list-item-action-text>
                                        Duração: <strong>{{ item.duracao }}</strong> -
                                        Distância em Linha Reta: <strong> {{ item.distanciaLinhaReta }}</strong> -
                                        Distância OLC: <strong>{{ item.olcKm }}</strong> -
                                        Pontuação OLC: <strong>{{ item.pontuacaoOlc }}</strong>
                                    </v-list-item-action-text>
                                </v-list-item-content>
                            </v-list-item>
                        </template>
                    </v-list>
                </v-card>
            </v-col>
        </v-row>
    </v-container>

</template>
<script>
import { obterListaVoos } from "@/services/voos"

export default {
    data() {
        return {
            lista: [],
            drawer: null
        }
    },
    methods: {
        obter() {
            obterListaVoos().then(res => {
                this.lista = res.data
            }).catch((err) => {
                console.log("err", err)
            })
        },
        verVoo(link) {
            const linkLimpo = link.replace('/flight/', '')
            this.$router.push(`/voo/${linkLimpo}`)
            
        }
    },
    mounted() {
        this.obter()
    }
}

</script>