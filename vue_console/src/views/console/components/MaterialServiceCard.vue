<template>
  <v-card
    v-bind="$attrs"
    :class="classes"
    class="v-card--material--service pa-3"
   
  >
    <div class="d-flex grow flex-wrap">
      <v-sheet
        :color="color"
        width="100%"
        elevation="6"
        class="pa-7 v-card--material__heading header"
        dark
        v-ripple
        @click="show"
      >
        <div
          class="header-left"
          v-text="service.Info.ServiceId"
        />
        <div
          class="header-right"
          v-text="service.Info.ServiceType"
        />
      </v-sheet>
    </div>
    <div class="baseinfo">
      <base-v-divider size="small" orientation="left">{{$t('base')}}</base-v-divider>
      <div class="baseinfo-centext">
          <v-row>
            <v-col cols="2" class="baseinfo-centext_name">{{$t('serviceid')}}</v-col>
            <v-col cols="4" class="baseinfo-centext_value">{{service.Info.ServiceId}}</v-col>
            <v-col cols="2" class="baseinfo-centext_name">{{$t('servicetype')}}</v-col>
            <v-col cols="4" class="baseinfo-centext_value">{{service.Info.ServiceType}}</v-col>
            <v-col cols="2" class="baseinfo-centext_name">{{$t('servicetag')}}</v-col>
            <v-col cols="4" class="baseinfo-centext_value">{{service.Info.ServiceTag}}</v-col>
            <v-col cols="2" class="baseinfo-centext_name">{{$t('version')}}</v-col>
            <v-col cols="4" class="baseinfo-centext_value">{{service.Info.ServiceVersion}}</v-col>
            <v-col cols="2" class="baseinfo-centext_name">{{$t('servicecategory')}}</v-col>
            <v-col cols="4" class="baseinfo-centext_value">{{service.Info.ServiceCategory}}</v-col>
            <v-col cols="6"></v-col>

            <v-col cols="2" class="baseinfo-centext_name">{{$t('processid')}}</v-col>
            <v-col cols="4" class="baseinfo-centext_value">{{service.Info.Pid}}</v-col>
            <v-col cols="2" class="baseinfo-centext_name">{{$t('processname')}}</v-col>
            <v-col cols="4" class="baseinfo-centext_value">{{service.Info.Pname}}</v-col>
          </v-row>
      </div>
    </div>
    <v-expand-transition>
      <div v-show="ishow">
        <div class="setting" v-if="service.Info.Setting && Object.keys(service.Info.Setting).length > 0">
          <base-v-divider size="small" orientation="left">{{$t('servicesetting')}}</base-v-divider>
          <div class="setting-centext">
            <v-row no-gutters>
              <v-col cols="12" sm="6" v-for="(item,name) in service.Info.Setting" :key="name">
                <v-text-field v-if="!item.IsWrite"
                  class="setting-centext_field"
                  :value="item.Data"
                  :label="name"
                  disabled
                  dense
                  outlined
                  readonly
                ></v-text-field>
                <v-text-field v-else
                  class="setting-centext_field"
                  :value="item.Data"
                  :label="name"
                  dense
                  outlined
                  readonly
                  :append-icon="'mdi-content-save'"
                  @click:append="modifyServiceSetting(item.Data)"
                ></v-text-field>
              </v-col>
            </v-row>
          </div>
        </div>
        <div v-if="service.Info.ModuleMonitor && Object.keys(service.Info.ModuleMonitor).length > 0" class="modules">
          <base-v-divider size="small" orientation="left">{{$t('servicemodules')}}</base-v-divider>
          <div class="modules-centext">
            <v-tabs v-model="tab" height="35px">
              <v-tab class="text-capitalize" v-for="(item,index) in service.Info.ModuleMonitor" :key="index" :href="'#tab-' + index">{{index}}</v-tab>
            </v-tabs>

            <v-tabs-items class="modules-centext-items" v-model="tab">
              <v-tab-item
                v-for="(item,index) in service.Info.ModuleMonitor" :key="index" :value="'tab-' + index"
              >
                <template v-if="item.Setting && Object.keys(item.Setting).length > 0">
                  <v-row no-gutters>
                    <v-col cols="12" sm="6" v-for="(item,name) in item.Setting" :key="name">
                      <v-text-field v-if="!item.IsWrite"
                        class="setting-centext_field text--disabled"
                        :value="item.Data"
                        :label="name"
                        disabled
                        dense
                        outlined
                        readonly
                      ></v-text-field>
                      <v-text-field v-else
                        class="setting-centext_field"
                        v-model="item.NewData"
                        :value="item.Data"
                        :label="name"
                        dense
                        outlined
                        :append-icon="'mdi-content-save'"
                        @click:append="modifyModuleSetting(index,name,item.Data,item.NewData)"
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </template>
              <template v-else>
                <div class="modules-centext-nodes">
                  {{$t('noconfiguration')}}
                </div>
              </template>
              </v-tab-item>
            </v-tabs-items>
          </div>
        </div>
        <div class="monitor">
            <base-v-divider size="small" orientation="left">{{$t('servicemonitor')}}</base-v-divider>
            <div class="monitor-centext">
            <v-tabs v-model="tab1" height="35px">
              <v-tab href="#tab1-0">{{$t('servicepreweight')}}</v-tab>
              <v-tab href="#tab1-1">{{$t('servicegoroutine')}}</v-tab>
              <v-tab href="#tab1-2">{{$t('cpu')}}</v-tab>
              <v-tab href="#tab1-3">{{$t('memory')}}</v-tab>
            </v-tabs>
            <v-tabs-items v-model="tab1">
              <v-tab-item value="tab1-0">
                  <LineChart height="300" :keys="service.Monitor.Keys" :values="service.Monitor.PreWeight"></LineChart>
              </v-tab-item>
              <v-tab-item value="tab1-1">
                  <LineChart  height="300" :keys="service.Monitor.Keys" :values="service.Monitor.Goroutine"></LineChart>
              </v-tab-item>
              <v-tab-item value="tab1-2">
                  <LineChart  height="300" :keys="service.Monitor.Keys" :values="service.Monitor.Cpu"></LineChart>
              </v-tab-item>
              <v-tab-item value="tab1-3">
                  <LineChart  height="300" :keys="service.Monitor.Keys" :values="service.Monitor.Memory"></LineChart>
              </v-tab-item>
            </v-tabs-items>
          </div>
        </div>
      </div>
    </v-expand-transition>
  </v-card>
</template>

<script>
import {setservicesetting,setmodulesetting} from '@/api/console'
import {LineChart} from "@/components/echarts"
export default {
  name: 'MaterialServiceCard',
  components:{
    LineChart
  },
  props: {
    color: {
      type: String,
      default: 'success',
    },
    service:{
      type: Object,
      default: undefined,
    }
  },
  computed: {
    classes () {
      return {
        'v-card--material--has-heading': this.hasHeading,
      }
    },
    hasHeading () {
      return Boolean(this.$slots.heading || this.title || this.icon)
    },
    hasAltHeading () {
      return Boolean(this.$slots.heading || (this.title && this.icon))
    },
  },
  data :function(){
    return {
      ishow:false,
      tab:null,
      tab1:null,
    }
  },
  methods:{
    show(){
      this.ishow = !this.ishow
    },
    modifyServiceSetting(key,oldvalue,newvalue){
        if ( oldvalue !=  newvalue){
          setservicesetting({sid: this.service.ServiceId, key: key,value:newvalue }).then(response => {
              const {message} = response
              this.$message.success(message)
            }).catch(error => {
              this.$message.error(error.message)
            })
        }
    },
    modifyModuleSetting(name,key,oldvalue,newvalue){
      if ( oldvalue !=  newvalue){
         setmodulesetting({sid: this.service.ServiceId, module: name, key: key,value:newvalue }).then(response => {
            const {message} = response
            this.$message.success(message)
          }).catch(error => {
            this.$message.error(error.message)
          })
      }
    }
  }
}
</script>

<style lang="sass">
  $title-color:#757575
  .v-card--material--service
    border-radius: 6px
    margin-top: 30px
    margin-bottom: 15px
    .header
      display: flex
      justify-content: space-between
      .header-left
        font-size: 2em
      .header-right
        font-size: 1em
        line-height: 50px
        text-align: center
        color: #F5F5F5
    .baseinfo
      width: 100%
      color: $title-color
      margin-bottom: 15px
      &-centext
        padding-left: 20px
        &_name
          color: #616161
          font-size: 0.8em
          height: 15px
          line-height: 15px
        &_value
          color: #BDBDBD
          font-size: 0.8em
          height: 15px
          line-height: 15px
    .setting
      width: 100%
      color: $title-color
      &-centext
        padding: 0px 20px 0px 20px
        &_field
          padding-right: 10px
          height: 50px
    .modules
      width: 100%
      color: $title-color
      &-centext
        padding: 0px 20px 0px 20px
        &-nodes
          height: 80px
          line-height: 80px
          vertical-align: baseline
          text-align: center
          color: #757575
          font-size: 0.7em
        &-items
          padding-top:10px
    .monitor
      width: 100%
      color: $title-color
      &-centext
        padding: 0px 20px 0px 20px
</style>
