<script setup>
const add_hub = async () => {
    if (pool_name === "") {
        return
    }
    const body = {
        type: "ADD_POOL_NAME",
        data: pool_name.value
    }
    const { data, error, refresh } = await useFetch("http://localhost:8080/addpool", {
        method: "POST",
        body: body,
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
        }
    })
    await props.refreshPools()
    console.log(data.value, error)
}
const pool_name = ref("")

const props = defineProps({
    refreshPools: Function
})

</script>

<template>
    <div>
        <div>
            <button class="bg-green-300 p-4 ring-2 ring-zinc-400 rounded-lg m-2" @click="add_hub">Add Random Hub</button>
            <br>
            <div class="m-2 p-2">
                <label for="pool_name">Enter Pool Name: </label>
                <input class="ring-2 m-2" name="pool_name" v-model="pool_name" />
            </div>
        </div>
    </div>
</template>