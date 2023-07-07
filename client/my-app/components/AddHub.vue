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
            <button @click="add_hub">Add Random Hub</button>
            <br>
            <div>
                <label for="pool_name">Enter Pool Name: </label>
                <input name="pool_name" v-model="pool_name" />
            </div>
        </div>
    </div>
</template>