<template>
    <div class="grid">
        <div class="col-12">
            <div class="card">
                <h5>Prometheus 实例</h5>
                <DataTable
                    v-model:selection="selectedInstance"
                    selectionMode="single"
                    :value="instances"
                    :paginator="true"
                    class="p-datatable-gridlines"
                    :rows="20"
                    dataKey="id"
                    :rowHover="true"
                    v-model:filters="filters"
                    filterDisplay="menu"
                    :loading="loading"
                    responsiveLayout="scroll"
                    :globalFilterFields="['AppName', 'Endpoint']"
                >
                    <template #header>
                        <div class="flex justify-content-between flex-column sm:flex-row">
                            <Button type="button" icon="pi pi-filter-slash" label="Clear" class="p-button-outlined mb-2" @click="clearFilter()" />
                            <span class="p-input-icon-left mb-2">
                                <i class="pi pi-search" />
                                <InputText v-model="filters['global'].value" placeholder="Keyword Search" style="width: 100%" />
                            </span>
                        </div>
                    </template>
                    <template #empty> No prometheus instances found. </template>
                    <template #loading> Loading prometheus instances. Please wait. </template>


                    <!-- columes -->
                    <Column field="AppName" header="实例" style="min-width: 12rem">
                        <template #body="{ data }">
                            {{ data.AppName }}
                        </template>
                    </Column>
                    <Column field="Endpoint" header="Endpoint" style="min-width: 12rem">
                        <template #body="{ data }">
                            <span style="margin-left: 0.5em; vertical-align: middle" class="image-text">{{ data.Endpoint }}</span>
                        </template>
                    </Column>
                    <Column field="HealthStatus" header="健康状态" dataType="boolean" bodyClass="text-center" style="min-width: 8rem">
                        <template #body="{ data }">
                            <i class="pi" :class="{ 'text-green-500 pi-check-circle': data.HealthStatus, 'text-pink-500 pi-times-circle': !data.HealthStatus }"></i>
                        </template>
                        <!-- <template #filter="{ filterModel }">
                            <TriStateCheckbox v-model="filterModel.value" />
                        </template> -->
                    </Column>
                    <Column selectionMode="single" header="操作" style="min-width: 10rem">
                        <template #body="{data}">
                            <div class="flex gap-2">
                                <Button type="button" label="reload" severity="success"  :loading="reloadLoading" @click="(event) => reload(event , data.Endpoint)" class="bg-green-500" raised/>
                                <Button type="button" label="restart"  @click="() => toast.add({severity: 'success'})" raised disabled/>
                            </div>
                        </template>
                    </Column>
                </DataTable>
            </div>
        </div>
    </div>
</template>

<script setup>
import { FilterMatchMode, FilterOperator } from 'primevue/api';
import CustomerService from '@/service/CustomerService';
import ProductService from '@/service/ProductService';
import PromService from '@/service/PromService';
import { ref, onBeforeMount } from 'vue';
import { useToast } from 'primevue/usetoast';

// data
const instances = ref([])
const expandedRows = ref([]);
const filters = ref(null);
const loading = ref(null);
let id = ref(0);
const selectedInstance = ref();
const reloadLoading = ref(false);
const toast = useToast();


// methods

function reload(event, endpoint) {
    reloadLoading.value = true;

    const url = "/proxy?host=" + endpoint + "&path=/-/reload"
    console.log(endpoint)
    toast.add({
        severity: 'success',
        summary: endpoint+ 'reload success',
        life: 3000
    });
    console.log(import.meta.env.PROD)
    console.log()
    console.log("base: "+import.meta.env.BASE_URL);
    // fetch(url, {
    //     method: 'GET',
    // }).then(
    //     (response) => {
    //         if (response.ok) {
    //             console.log(event)
    //         }
    //     }
    // )
    setTimeout(() => {
        reloadLoading.value = false;
    }, 1000);
    // reloadLoading.value = false;
}


const customerService = new CustomerService();
const productService = new ProductService();
const promService = new PromService();

onBeforeMount(() => {
    // promService.getInstances().then(
    //     (data) => {
    //         instances.value = data.instances;
    //         instances.value.forEach((v) => {
    //             v.Endpoint = `${v.ServiceAddress}:${v.ServicePort}`
    //             v.AppName = v.ServiceTags.filter(
    //                 (tag) => {
    //                     return tag.startsWith('app')
    //                 }
    //             )[0];
    //             v.HealthStatus = true;
    //             v.id = id.value++;
    //         });

    //         console.log(instances.value)
    //         loading.value = false;
    //     });
    console.log(import.meta.env.PROD)

    promService.getInstances().then(
        (data) => {
            instances.value = data;
            instances.value.forEach((v) => {
                v.Endpoint = `${v.ServiceAddress}:${v.ServicePort}`
                v.AppName = v.ServiceTags.filter(
                    (tag) => {
                        return tag.startsWith('app')
                    }
                )[0];
                v.HealthStatus = true;
                v.id = id.value++;
            });

            console.log(instances.value)
            console.log(import.meta.env.PROD)
            loading.value = false;
        });
    
    initFilters();
});

const initFilters = () => {
    filters.value = {
        global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    };
};

const clearFilter = () => {
    initFilters();
};
const expandAll = () => {
    expandedRows.value = products.value.filter((p) => p.id);
};
const collapseAll = () => {
    expandedRows.value = null;
};
const formatCurrency = (value) => {
    return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
};

const formatDate = (value) => {
    return value.toLocaleDateString('en-US', {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric'
    });
};
</script>



<style scoped lang="scss">
::v-deep(.p-datatable-frozen-tbody) {
    font-weight: bold;
}

::v-deep(.p-datatable-scrollable .p-frozen-column) {
    font-weight: bold;
}
</style>
