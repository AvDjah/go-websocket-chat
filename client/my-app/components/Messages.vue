<template>
    <div>
        Messages from client: {{ current_pool === null ? "Not Connected" : current_pool }}
        <div>
            <div v-for="message of message_from_websocket">
                <div class="message border-2 p-6 ">
                    <div style="width: 10px; height: 20px; border: 1px solid black; display: inline-block; margin-right: 10px;"
                        :class="{ green: is_green(message.from), red: !is_green(message.from) }">
                    </div>
                    <div>
                        {{ message.message }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>

const message_from_websocket = useState("message_from_websocket")

const is_green = (from) => {
    return from === 0
}

const current_pool = useState("current_pool")

</script>

<style scoped>
.red {
    background-color: red;
}

.green {
    background-color: green;
}

span {
    background-color: red;
    width: 20px;
    height: 20px;
}

.message {
    display: flex;
    margin: 5px;
}
</style>