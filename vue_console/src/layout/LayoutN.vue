<template>
    <v-app>
        <v-main>
            <v-row justify="start" class="topbar">
                <v-spacer></v-spacer>
                <SelectLanguge/>
                <Settings/>
            </v-row>
            <v-container
            class="fill-height"
            fluid
            >
                <v-row
                    align="center"
                    justify="center"
                >
                    <v-col
                        cols="12"
                        sm="8"
                        md="4"
                    >
                    <transition name="fade-transform" mode="out-in">
                        <router-view></router-view>
                    </transition>
                    </v-col>
                </v-row>
            </v-container>
        </v-main>
    </v-app>
</template>

<script>
import {SelectLanguge,Settings} from "./components"

export default {
    name:"LayoutN",
    components:{
        SelectLanguge,
        Settings,
    }
}
</script>

<style>
    .topbar{
        width: 100%;
        padding: 10px;
    }
</style>