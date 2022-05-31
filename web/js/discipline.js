// const Modal = {
//     data() {
//         return {
//             disciplineName: ''
//         }
//     }
// }
//
// Vue.createApp(Modal).mount('#app')

Vue.createApp({
    data() {
        return {
            disciplineName: ''
        }
    }
    }
).mount('#app')


// Vue.createApp({
//     data() {
//         return {
//         disciplines: [],
//         groups: [
//             { name: 'A-13-18' },
//             { name: 'A-05-19' },
//             { name: 'А-13-20' },
//             { name: 'Э-13-20'},
//             { name: 'ИЭ-01-20' },
//             { name: 'ИЭ-02-20' }
//         ],
//         selectGroup: ''
//         }
//     },
//     methods: {
//         addDisciplines() {
//
//         },
//         created() {
//             this.$https.get(`http://localhost:8080/api/discipline`, this.disciplines).then(resp => {
//                 console.log('resp:', resp)
//                 this.disciplines = resp.body
//             }).catch((resp) => {
//                 console.log('err:', resp)
//             })
//         }
//
//     },
//
// }).mount('#v-model-select')
