<template>
    <div class="menu-butt">
       <v-menu
            :close-on-content-click="false"
            :nudge-width="200"
            offset-x
            offset-y
            transition="scroll-x-transition"
        >
        <template v-slot:activator="{ on, attrs }">
            <img
                class="avatar-butt"
                v-bind="attrs"
                v-on="on"
                :src="userinfo.Headurl?userinfo.Headurl:require('@/assets/defavatar.jpg')"
            >
        </template>

        <v-card
            width="320px"
        >
            <div class="avatar-title">
                <div class="avatar-title-bg"></div>
                <v-row class="avatar-title-center" justify="center">
                    <v-col cols="2" align-self="center">
                        <v-avatar>
                            <img :src="userinfo.Headurl?userinfo.Headurl:require('@/assets/defavatar.jpg')">
                        </v-avatar>
                    </v-col>
                    <v-col cols="9" align-self="center">
                        <div class="avatar-title-text">
                            <p class="p-name">{{userinfo.NickName}}</p>
                            <p class="p-eamil">{{userinfo.Email}}</p>
                        </div>
                    </v-col>
                </v-row>
            </div>

            <v-divider></v-divider>

            <v-list>
                <v-list-item>
                    <v-list-item-action>
                    <v-btn class="text-capitalize" text>{{$t('common.profile')}}</v-btn>
                    </v-list-item-action>
                </v-list-item>
                <v-list-item>
                    <v-list-item-action>
                    <v-btn text @click="logout">{{$t('common.LoginOut')}}</v-btn>
                    </v-list-item-action>
                </v-list-item>
            </v-list>
        </v-card>
        </v-menu>
    </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
    name : "AvatarMenu",
    props: {
        
    },
    computed: {
        ...mapGetters([
        'currtagroute',
        'userinfo'
        ]),
    },
    watch: {
        
    },
    data() {
        return {
        }
    },
    methods:{
        logout(){
            this.$store.dispatch('user/loginout')
                .then(() => {
                    this.$router.push({ path:'/base/login' })
                })
                .catch((error) => {
                    this.$message.error(error.message)
                }
            )
        }
    },
}
</script>

<style lang="scss" scoped>
    .menu-butt{
        width: 54px;
        height: 46px;
        display: flex;
        align-items: center;
    }
    .avatar-butt{
        width: 46px;
        height: 46px;
        border-radius:100%;
    }
    .avatar-title{
        width: 100%;
        height: 111px;
        background-color: #30b1ff;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .avatar-title-bg{
        width: 100%;
        height: 111px;
        opacity: .05;
        background-image: url(~@/assets/media/city3.bade0f80.jpg);
        position: absolute;
        z-index: 8;
    }
    .avatar-title-center{
        width: 100%;
        position: absolute;
        z-index: 10;
    }
    .avatar-title-text{
       text-align: start;
       
       .p-name{
            font-size: 1.2em;
            line-height: 25px;
            margin: 0;
            vertical-align:middle;
            color: #ffffff;
        }
        .p-eamil{
            margin: 0;
            font-size: 1em;
            line-height: 25px;
            vertical-align: middle;
            color: #EEEEEE;
       }
    }
</style>