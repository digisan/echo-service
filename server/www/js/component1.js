import { getEmitter } from './mitt.js'
import { fetch_get_json } from './fetch.js'
import { local_ws } from './ws.js'

let emitter = getEmitter();

export default {
    name: 'Test',

    setup() {
        const title = "Hello";
        let mypen = Vue.ref("");
        let imgsrc = Vue.ref("");
        let timer_str = Vue.ref("");
        let ws_str = Vue.ref("");

        // listen to an event
        emitter.on('from_app1', e => {
            mypen.value = e;
            console.log('app received:', e);
        });

        // listen to all events
        // emitter.on('*', (type, e) => console.log(type, e));

        // fire an event
        // emitter.emit('foo', { a: 'b' })

        function fire() {

            let cData = fetch_get_json('https://yesno.wtf/api');

            // 'async function' return channel             
            const fnFetchValue = async () => {
                const data = await cData;
                emitter.emit('from_app', data.answer);
                console.log(data.answer);
                imgsrc.value = data.image;
            };
            // 'async function' return channel
            let cOut = fnFetchValue();
            console.log(`com1 result is ${cOut}`)

            // emitter.emit('from_app', mypen.value);
            // console.log(mypen.value);
        }

        // timer sample
        let myTimer = setInterval(
            () => { timer_str.value = (new Date()).toLocaleTimeString(); },
            1000,
        )
        // clearInterval(myTimer);

        // web socket sample
        let ws = local_ws("ws/test"); // registered in server reg_api
        ws.onopen = function () {
            console.log('Connected')
        }
        ws.onmessage = function (evt) {
            ws_str.value = evt.data;
        }
        let idx = 0;
        // trigger server to send back message, handle it in 'onmessage'
        setTimeout(() => { ws.send('Hello, Server! ' + idx.toString()); idx++ }, 1000);

        return {
            title,
            mypen,
            fire,
            imgsrc,
            timer_str,
            ws_str,
        };
    },

    template: `      
        <h1>{{title}} | {{timer_str}} | {{ws_str}} | {{mypen}} | {{imgsrc}}</h1>
        <input v-model="mypen" placeholder="input">
        <button class="mybutton" @click="fire"></button>   
        <img :src="imgsrc" alt="YES/NO IMAGE" width="320" height="240" />   
    `,
};