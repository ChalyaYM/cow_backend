<template>
<div class="lac-title">Лактации</div>    
<div v-if="!isLoading">
    <button @click="isTable=true;isChart=false" 
    class="table-chart"
    :class="{'active-btn':isTable}">
    Таблица</button>
    <button 
    @click="isTable=false;isChart=true" 
    class="table-chart"
    :class="{'active-btn':isChart}">
    График</button>
    <div class="parent-table" v-if="isTable">
        <table class="lac-table">
            <thead>
                <tr class="lac-header">
                    <th>Номер лактации</th>
                    <!-- <th>Кратность осеменения</th> -->
                    <th>Дата осеменения к текущей лактации</th>
                    <!-- <th>Количество телят</th> -->
                    <!-- <th>Аборт</th> -->
                    <th>Дата отела</th>
                    <th>Количество дойных дней</th>
                    <th>Удой за всю лактацию, кг</th>
                    <th>Содержание жира за всю лактацию, %</th>
                    <th>Выход молочного жира за всю лактацию, кг</th>
                    <th>Содержание белка за всю лактацию, %</th>
                    <th>Выход молочного белка за всю лактацию, кг</th>
                    <th>Удой за 305 дней, кг</th>
                    <th>Содержание жира за 305 дней лактации, %</th>
                    <th>Выход молочного жира за 305 дней, кг</th>
                    <th>Содержание белка за 305 дней лактации, %</th>
                    <th>Выход молочного белка за 305 дней, кг</th>
                    <th>Дата запуска</th>
                    <th>Результат отёла</th>
                    <th>Сервис-период, дней</th>
                </tr>
            </thead>
            <tbody class="lac-tablebody">
                <tr v-for="lact in cow_info" :key="lact.Number">
                    <td>{{ lact.Number || 'Нет информации'}}</td>
                    <!-- <td>{{ lact.InsemenationNum || 'Нет информации'}}</td> -->
                    <td v-if="lact.InsemenationDate">{{ dateConverter(lact.InsemenationDate) }}</td><td v-else>Нет информации</td>
                    <!-- <td v-if="lact.CalvingCount">{{ lact.CalvingCount || 'Нет информации'}}</td><td v-else>Нет информации</td> -->
                    <!-- <td v-if="lact.Abort===true || lact.Abort===false">{{ yesNo(lact.Abort) }}</td><td v-else>Нет информации</td> -->
                    <td v-if="lact.CalvingDate">{{ dateConverter(lact.CalvingDate) }}</td><td v-else>Нет информации</td>
                    <td>{{ lact.Days || 'Нет информации'}}</td>
                    <td v-if="lact.MilkAll">{{ Math.floor(lact.MilkAll) }}</td><td v-else>Нет информации</td>
                    <td></td>
                    <td v-if="lact.FatAll">{{ Math.floor(lact.FatAll) }}</td><td v-else>Нет информации</td>
                    <td></td>
                    <td v-if="lact.ProteinAll">{{ Math.floor(lact.ProteinAll) }}</td><td v-else>Нет информации</td>
                    <td v-if="lact.Milk305">{{ Math.floor(lact.Milk305) }}</td><td v-else>Нет информации</td>
                    <td></td>
                    <td v-if="lact.Fat305">{{ Math.floor(lact.Fat305) }}</td><td v-else>Нет информации</td>
                    <td></td>
                    <td v-if="lact.Protein305">{{ Math.floor(lact.Protein305) }}</td><td v-else>Нет информации</td>
                    <td></td>
                    <td></td>
                    <td>{{ lact.ServicePeriod || 'Нет информации'}}</td>
                </tr>
            </tbody>
        </table>
    </div>

    <div v-if="isChart">
        <div class="chart-flex">
            <div class="chart-param">Показатель: </div>
            <select v-model="param_milking" class="select-param">
                <option value="MilkAll">Суммарный удой, кг</option>
                <option value="Milk305">Суммарный удой за 305 дней, кг</option>
                <option value="FatAll">Суммарный жир, кг</option>
                <option value="Fat305">Суммарный жир за 305 дней, кг</option>
                <option value="ProteinAll">Суммарный белок, кг</option>
                <option value="Protein305">Суммарный белок за 305 дней, кг</option>
                <option value="MilkDaily">Удой среднесуточный, кг</option>
                <option value="FatDaily">Жир среднесуточный, кг</option>
                <option value="ProteinDaily">Белок среднесуточный, кг</option>
                <option value="Days">Количество дойных дней</option>
                <option value="ServicePeriod">Длительность сервис-периода</option>
            </select>
        </div>

        <apexchart id="ControlMilking" width="600" type="bar" :options="options" :series="series"></apexchart>
    </div>
</div>
<div v-else>Идёт загрузка...</div>
</template>

<script>
export default {
    data() {
        return {
            cow_info: [],
            isTable: true,
            isChart: false,
            options: {
                chart: {
                    id: 'ControlMilking'
                },
                xaxis: {
                    categories: []
                },
                tooltip: {
                    enabled: false,
                }
            },
            series: [],
            param_milking: 'MilkAll',

            isLoading: false,
        }
    },
    async created() {
        this.isLoading = true;
        let mass_route = this.$route.path.split('/');
        let cow_id = mass_route[2];
        let response = await fetch(`/api/cows/${cow_id}/lactations`);
        let result = await response.json();
        this.cow_info = result;

        let serie = {name:'Удой полный',data: []};
        for (let i = 0; i < this.cow_info.length; i++) {
            this.cow_info[i].MilkDaily = Math.round(this.cow_info[i].MilkAll / this.cow_info[i].Days);
            this.cow_info[i].FatDaily = Math.round(this.cow_info[i].FatAll / this.cow_info[i].Days);
            this.cow_info[i].ProteinDaily = Math.round(this.cow_info[i].ProteinAll / this.cow_info[i].Days);

            this.options.xaxis.categories.push('Лактация ' + this.cow_info[i].Number);
            serie.data.push(this.cow_info[i].MilkAll);
        }
        this.series.push(serie);
        this.isLoading = false;
    },
    methods: {
        addParam(obj, arr, param) {
            let serie = {
                name: `${param}`,
                data: []
            };
            for (let i = 0; i < obj.length; i++) {
                serie.data.push(Math.floor(obj[i][param]));
            }
            arr.push(serie);
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        },
        yesNo(param) {
            if(param) return 'Да';
            else return 'Нет';
        }
    },
    watch: {
        param_milking(new_value) {
            this.series = [];
            this.addParam(this.cow_info, this.series, new_value);
        }
    }
}
</script>

<style scoped>
.lac-title {
    font-size: 130%;
    color: rgb(37, 0, 132);
    padding-bottom: 30px;
    width: max-content;
}

.parent-table {
    width: 49vw;
    overflow-x: auto;
}

.lac-table {
    margin-bottom: 30px;
    text-align: left;
}

th {
    font-weight: normal;
}

td {
    width: auto;
    min-width: 130px;
    padding-right: 7px;
}

.lac-header {
    color: grey;
}

.lac-header th {
    padding-right: 30px;
    padding-bottom: 15px;
}

.lac-tablebody {
    text-align: left;
}

.parent-table::-webkit-scrollbar {
    height: 12px;
}

.parent-table::-webkit-scrollbar-track {
    background: rgb(241, 241, 241);
}

.parent-table::-webkit-scrollbar-thumb {
    background-color: rgb(183, 183, 183);
    border-radius: 20px;
    border: 3px solid rgb(241, 241, 241);
}

.table-chart {
    border: 1px solid rgb(37, 0, 132);
    background-color: white;
    color: rgb(37, 0, 132);
    padding: 10px 0;
    margin: 20px 0;
    width: 100px;
    border-radius: 10px;
    transition: 0.3s;
    cursor: pointer;
    margin-right: 15px;
    transition: 0.3s;
}

.active-btn {
    background-color: rgb(219, 214, 239);
}

.chart-param {
    width: max-content;
    padding: 6px 4px;
    margin-bottom: 3px;
    font-size: 110%;
}

.chart-flex {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 350px;
}

.select-param {
    padding: 5px;
    border: 1px solid rgb(37, 0, 132);
    border-radius: 3px;
}
</style>