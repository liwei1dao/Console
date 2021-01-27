<template>
    <div class="menu-butt">
       <v-menu
            offset-x
            offset-y
            :close-on-content-click="false"
            transition="scroll-x-transition"
        >
        <template v-slot:activator="{ on, attrs }">
            <div class="setting-butt">
                <v-icon
                    v-bind="attrs"
                    v-on="on"
                    size="35px"
                >mdi-cog</v-icon>
            </div>
   
        </template>

        <v-card
            width="320px"
        >
            <div class="setting-title">
                <div class="setting-title-bg"></div>
                <div class="setting-title-center">{{$t('common.Setting')}}</div>
            </div>
            <v-divider></v-divider>
            <v-list>
                <v-subheader>{{$t('common.SysSetting')}}</v-subheader>
                <v-divider></v-divider>
                <v-list-item>
                    <v-list-item-title>{{$t('common.SysSetting-Dark')}}</v-list-item-title>
                    <v-list-item-action><v-switch :input-value="dark"  v-on:change="darkchange"></v-switch></v-list-item-action>
                </v-list-item>
            </v-list>
        </v-card>
        </v-menu>
    </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
    name:"Settings",
    computed: {
        ...mapGetters([
            'dark',
        ]),
    },
    data() {
        return {

        }
    },
    methods:{
        darkchange(value){
            console.log("设置深度:",value)
            this.$store.dispatch('app/setcurrdark', value)
            this.$vuetify.theme.dark = value
        }
    }
}
</script>

<style lang="scss" scoped>
    .menu-butt{
        width: 54px;
        height: 46px;
        display: flex;
        align-items: center;
    }
    .setting-butt{
        width: 46px;
        height: 46px;
        background-color: #E0E0E0;
        border-radius:100%;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .setting-title{
        width: 100%;
        height: 65px;
        background-image: linear-gradient(135deg,#667eea,#764ba2);
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .setting-title-bg{
        width: 100%;
        height: 65px;
        opacity: .2;
        background-image: url(~@/assets/media/abstract4.d73aa981.jpg);
        position: absolute;
        z-index: 8;
    }
    .setting-title-center{
        width: 100%;
        position: absolute;
        z-index: 10;
        display: flex;
        align-items: center;
        justify-content: center;
        color: #fff;
    }
</style>