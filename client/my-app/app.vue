<template>
  <div>
    <div class="text-3xl text-white font-bold bg-red-400 border-2 p-4 m-4">Socks 🧦</div>
    <WebsocketStatus :reconnect-websocket="() => { connect_to_websocket() }">
    </WebsocketStatus>
    <div>
      <input v-model="text" @submit="send_message_to_websocket">
    </div>
    <div>
      <button @click="send_message_to_websocket">Send To Websocket</button>
    </div>
    <div>
      <Messages></Messages>
    </div>
    <hr style="height: 10px;">
    <div>
      <PoolList :get-pools="get_pool_list" :send-pool-id="send_pool_id"></PoolList>
    </div>
    <hr style="height: 10px;">
    <AddHub :refresh-pools="get_pool_list"></AddHub>
  </div>
</template>


<script setup>

let webSocket

const text = ref("")


const pool_list = useState("pool_list", () => [])

const websocket_status = useState("websocket_status", () => false)

const message_from_websocket = useState("message_from_websocket", () => [])

const websocket_connected = ref(false)

const current_pool = useState("current_pool", () => null)



const get_pool_list = async (event) => {
  const { data, error, refresh } = await useFetch("http://localhost:8080/sendpools")
  pool_list.value = data.value
  console.log(data.value)
}

const connect_to_websocket = () => {
  webSocket = new WebSocket("ws://localhost:8080/ws")
  console.log(webSocket)


  // WEBSOCKET IS CONNECTED
  webSocket.onopen = async () => {
    websocket_connected.value = true
    console.log("Connected!!")
    websocket_status.value = true
    const { data, error, refresh } = await useFetch("http://localhost:8080/sendpools")
    pool_list.value = data.value
    console.log("Pool List Fetched")
  }

  // WEBSOCKET CLOSE DUE TO ERROR
  webSocket.onerror = (err) => {
    websocket_status.value = false
    console.log("Websocket Error")
    // webSocket.close()
  }

  // WEBSOCKET CLOSED SO RETRY
  webSocket.onclose = () => {
    websocket_status.value = false
    console.log("Trying Again")
    // setTimeout(connect_to_websocket, 1000)
  }

  // WEBSOCKET MESSAGE RECEIVED
  webSocket.onmessage = (event) => {
    console.log("Received: ", event.data)

    const json = JSON.parse(event.data)
    if (json.type === "POOL_ID_JOIN_RESULT") {
      if (json.data !== "-1") {
        current_pool.value = json.data
      } else {
        current_pool.value = "No Pool/ Pool Error"
      }
      return
    }

    message_from_websocket.value.push({
      message: json.data,
      from: 1
    })
  }
}

onMounted(() => {
  connect_to_websocket()
})

const send_pool_id = (event, id) => {
  console.log(id.pool_id)
  const resp = {
    type: "POOL_ID_JOIN",
    data: id.pool_id
  }
  webSocket.send(JSON.stringify(resp))
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

<style scoped>
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
</style>
