<template>
  <div>
    <div>
      <input v-model="text" @submit="send_message_to_websocket">
    </div>
    <div>
      <button @click="send_message_to_websocket">Send To Websocket</button>
    </div>
    <div v-if="websocket_connected.valueOf()" class="connecte1d">
      Websocket is connect
    </div>
    <div v-else class="disconnected">
      Disconnect : <button>Reconnect</button>
    </div>
    <div>
      Messages from client:
      <div>
        <div v-for="message of message_from_websocket.values()">
          <div class="message">
            <div style="width: 10px; height: 20px; border: 1px solid black; display: inline-block;"
              :class="{ green: is_green(message.from), red: !is_green(message.from) }">
            </div>
            <div>
              {{ message.message }}
            </div>
          </div>
        </div>
      </div>
    </div>
    <hr style="height: 10px;">
    <div>
      <button @click="get_pool_list">Get All Pools</button>
      <br /> Pool List:
      <li class="pool-id-item" v-for="pool of pool_list.values()" @click="(event) => send_pool_id(event, pool)">
        {{ pool.name }}
      </li>
    </div>
    <hr style="height: 10px;">
    <AddHub :refresh-pools="get_pool_list"></AddHub>
  </div>
</template>


<script setup>

let webSocket

const text = ref("")

const pool_list = ref([])

const message_from_websocket = ref([])

const websocket_connected = ref(false)


const send_pool_id = (event, id) => {
  console.log(id.pool_id)
  const resp = {
    type: "POOL_ID_JOIN",
    data: id.pool_id
  }
  webSocket.send(JSON.stringify(resp))
}


const get_pool_list = async (event) => {
  const { data, error, refresh } = await useFetch("http://localhost:8080/sendpools")
  pool_list.value = data.value
  console.log(data.value)
}

onMounted(() => {

  webSocket = new WebSocket("ws://localhost:8080/ws")
  console.log(webSocket)
  webSocket.onopen = async () => {
    websocket_connected.value = true
    console.log("Connected!!")
    const { data, error, refresh } = await useFetch("http://localhost:8080/sendpools")
    pool_list.value = data.value
    console.log("Pool List Fetched")
  }
  webSocket.onerror = (err) => {
    websocket_connected = false
    console.log(err)
  }
  webSocket.onmessage = (event) => {
    console.log("Received: ", event.data)
    message_from_websocket.value.push({
      message: event.data,
      from: 1
    })
  }
})

const is_green = (from) => {
  return from === 0
}

const send_message_to_websocket = (event) => {
  const resp = {
    type: "MESSAGE",
    data: text.value
  }
  if (text.value === "") {
    return
  }
  message_from_websocket.value.push({
    message: text.value,
    from: 0
  })
  text.value = ""
  webSocket.send(JSON.stringify(resp))
}


</script>

<style>
.green {
  background-color: green;
}

span {
  background-color: red;
  width: 20px;
  height: 20px;
}

input {
  border: 2px solid blue;
  padding: 10px;
}

button {
  background-color: aquamarine;
  padding: 10px;
  color: black;
}

button:active {
  background-color: black;
  color: white;
}


div {
  margin: 10px;
}

.connected {
  color: green;
}

.disconnected {
  color: red;
}

hr {
  position: relative;
  top: 20px;
  border: none;
  height: 12px;
  background: black;
  margin-bottom: 50px;
}

.pool-id-item {
  cursor: pointer;
}

.pool-id-item:active {
  background-color: black;
  color: white;
}

.message {
  display: flex;
  margin: 0px;
}

.red {
  background-color: red;
}
</style>