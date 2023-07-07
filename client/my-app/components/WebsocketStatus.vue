<template>
    <div class="m-4">
        <div v-if="websocket_status" class="bg-lime-600 text-white p-4 m-2 text-center border-white border-2 rounded-lg
    hover:-translate-y-1 hover:scale-110 transition ease-in-out duration-300 inline-block md:fixed top-5 right-5 font-mono
    static">
            WebSocket status: Connected
        </div>
        <div v-else class="bg-red-500 text-white p-4 m-2 text-center  border-white border-2 rounded-lg
    hover:-translate-y-0.5 hover:scale-105 transition ease-in-out duration-300 inline-block md:fixed top-5 right-5 font-mono
    ">
            WebSocket status: Disconnected
            <br>
            <div @click="props.reconnectWebsocket"
                class="bg-yellow-400 text-gray-800 inline-block text-center p-2 rounded-md hover:text-gray-400 cursor-pointer">
                Reconnect
            </div>
        </div>
    </div>
</template>


<script setup>

const websocket_status = useState("websocket_status")

const props = defineProps({
    reconnectWebsocket: Function
})


</script>