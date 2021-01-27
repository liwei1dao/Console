const getters = {
    appname: state => state.app.appname,
    currlanguage: state => state.app.currlanguage,
    dark: state => state.app.dark,
    routes: state => state.router.routes,
    currtagroute : state => state.router.currtagroute,
    token:state => state.user.token,
    userinfo: state => state.user.userinfo,
    projectinfo:state => state.console.projectinfo,
    hostinfo:state => state.console.hostinfo,
    cpuinfo:state => state.console.cpuinfo,
    memoryinfo:state => state.console.memoryinfo,
    hostmonitordata:state => state.console.hostmonitordata,
    clusterservice: state => state.console.clusterservice,
}
export default getters