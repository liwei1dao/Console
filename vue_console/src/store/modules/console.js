import { getprojectinfo,gethostinfo,getcpuinfo,getmemoryinfo,gethostmonitordata,getconsolecluster } from '@/api/console'

const state = {
    projectinfo:null,
    hostinfo:null,
    cpuinfo:null,
    memoryinfo:null,
    hostmonitordata:null,
    clusterservice:null
 }
 
 const mutations = {
    Set_projectinfo: (state,projectinfo) => {
        state.projectinfo = projectinfo
    },
    Set_hostinfo: (state,hostinfo) => {
        state.hostinfo = hostinfo
    },
    Set_cpuinfo: (state,cpuinfo) => {
        state.cpuinfo = cpuinfo
    },
    Set_memoryInfo: (state,memoryinfo) => {
        state.memoryinfo = memoryinfo
    },
    Set_hostmonitordata: (state,hostmonitordata) => {
        state.hostmonitordata = hostmonitordata
    },
    Set_clusterservice: (state,clusterservice) => {
        for (var k in clusterservice) {  
            for (var k0 in clusterservice[k].Info.Setting) {
                clusterservice[k].Info.Setting[k0].NewData = clusterservice[k].Setting[k0].Data
            }

            for (var k1 in clusterservice[k].Info.ModuleMonitor) {  
                for (var k2 in clusterservice[k].Info.ModuleMonitor[k1].Setting) {  
                    clusterservice[k].Info.ModuleMonitor[k1].Setting[k2].NewData = clusterservice[k].Info.ModuleMonitor[k1].Setting[k2].Data
                }
            }
        }
        state.clusterservice = clusterservice
    },
 }
 
 const actions = {
    getprojectinfo({ commit }) {
        return new Promise((resolve, reject) => {
            getprojectinfo().then(response => {
                const {code,data } = response
                if (code == 0){
                    commit('Set_projectinfo',data)
                    resolve(response)
                }else{
                    reject(response)
                }
            }).catch(error => {
                reject(error)
            })
        })
    },
    gethostinfo({ commit }) {
        return new Promise((resolve, reject) => {
            gethostinfo().then(response => {
                const {code,data } = response
                if (code == 0){
                    commit('Set_hostinfo',data)
                    resolve(response)
                }else{
                    reject(response)
                }
            }).catch(error => {
                reject(error)
            })
        })
    },
    getcpuinfo({ commit }) {
        return new Promise((resolve, reject) => {
            getcpuinfo().then(response => {
                const {code,data } = response
                if (code == 0){
                    commit('Set_cpuinfo',data)
                    resolve(response)
                }else{
                    reject(response)
                }
            }).catch(error => {
                reject(error)
            })
        })
    },
    getmemoryinfo({ commit }) {
        return new Promise((resolve, reject) => {
            getmemoryinfo().then(response => {
                const {code,data } = response
                if (code == 0){
                    commit('Set_memoryInfo',data)
                    resolve(response)
                }else{
                    reject(response)
                }
            }).catch(error => {
                reject(error)
            })
        })
    },
    gethostmonitordata({ commit }) {
        return new Promise((resolve, reject) => {
            gethostmonitordata().then(response => {
                const {code,data } = response
                if (code == 0){
                    commit('Set_hostmonitordata',data)
                    resolve(response)
                }else{
                    reject(response)
                }
            }).catch(error => {
                reject(error)
            })
        })
    },
    getconsolecluster({ commit }) {
        return new Promise((resolve, reject) => {
            getconsolecluster().then(response => {
                const {code,data } = response
                if (code == 0){
                    commit('Set_clusterservice',data)
                    resolve(response)
                }else{
                    reject(response)
                }
            }).catch(error => {
                reject(error)
            })
        })
    },
 }
 
 export default {
   namespaced: true,
   state,
   mutations,
   actions
 }