<template>
    <div class="searchbar">  
        <div class="searchbarinput" :class="{'searchbarinput-show':isshowinput}">
            <v-text-field
                label="Outlined"
                rounded
                single-line
                filled
                dense
            ></v-text-field>
        </div>
        <v-btn class="mx-2 searchbarbutt" fab small @click="searchbarClick">
            <v-icon dark>mdi-magnify</v-icon>
        </v-btn>
    </div>
</template>

<script>

export default {
    name:"SearchBar",
    data:function(){
        return {
            isshowinput:false
        }
    },
    methods:{
        searchbarClick(){
            this.isshowinput = !this.isshowinput
        }
    }
}
</script>

<style lang="scss" scoped>
    .searchbar{
        height: 50px;
        line-height: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
    }
    .searchbarbutt{
        right: 20px;
    }
    .searchbarinput{
        height: 30px;
        width: 0px;
        opacity: 0;
        margin-right: 0px;
        transition: width 0.3s;
    }
    .searchbarinput-show{
        height: 30px;
        width: 236px;
        opacity:1;
    }
</style>