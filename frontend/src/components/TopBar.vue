<template>
    <div class="topbar flex-topbar" id="topbar" :class="{'topbar-main-style': this.$route.path != '/'}">
        <a href="http://www.vsau.ru/"><img src="../img/logo.png" width="75" class="img-logo"></a><div class="logo" @click="$router.push('/')">GenMilk</div>
        <div class="topbar-links" id="topbar-links">
            <div class="point-of-bar" @click="$router.push('/animals')" v-if="token">Животные</div>
            <div class="point-of-bar" @click="$router.push('/analytics')" v-if="token">Статистика</div>
            <!-- <div class="point-of-bar" @click="$router.push('/gui')" v-if="token">Загрузка файлов</div> -->
            <div class="point-of-bar" @click="$router.push('/help')">Помощь</div>
            <div class="point-of-bar" @click="$router.push('/')">О проекте</div>
            <div class="point-of-bar" @click="$router.push('/login')" v-if="!token">Войти</div>
            <div class="point-of-bar" @click="logout" v-if="token">Выход</div>
        </div>
        <div id="date-update" class="date-update" :class="{'date-update-mainstyle': this.$route.path != '/'}">Дата обновления базы данных: {{ dateBD }}</div>
    </div>
    
</template>

<script>
export default {
    data() {
        return {
            dateBD: '',
        }
    },
    created() {
        window.addEventListener('scroll', this.handleScroll);
    },
    methods: {
        handleScroll() {
            if (this.$route.path == '/') {
                let topbar = document.getElementById('topbar');
                let datebd = document.getElementById('date-update');
                if(window.scrollY > 0) {
                    topbar.classList.add('topbar-main-style');
                    datebd.classList.add('date-update-mainstyle');
                }
                if(window.scrollY == 0) {
                    if (topbar.classList.contains('topbar-main-style')) {
                        topbar.classList.remove('topbar-main-style');
                        datebd.classList.remove('date-update-mainstyle');
                    }
                }
            }
            
        },
        dateConverter(date) {
            let arr = date.split('-');
            let result = '';
            result += arr[2]; result += '.';
            result += arr[1]; result += '.';
            result += arr[0];
            return result;
        },
        logout(){
            document.cookie = 'jwt=0; max-age=-1; path=/';
            location.reload();
        }
    },
    async mounted() {
        try {
            let response = await fetch('/api/updates')
            let result = await response.json();
            let datetime = result.Date;
            let date_arr = datetime.split('T');
            this.dateBD = this.dateConverter(date_arr[0]);
        } catch(err) {
            this.dateBD = '';
        }
    },
    computed: {
        token(){
            let arr = document.cookie.split(';');
            for (let i = 0; i < arr.length; i++) {
                if (arr[i].split('=')[0] == 'jwt') {
                    return true;
                }
            }
            return false;
        }
    }
}
</script>

<style>
    .topbar {
        height: 110px;
        width: 100%;
        background-color: none;
        transition: 0.4s;
        color: white;
        z-index: 31;
    }

    .img-logo {
        align-self: flex-start;
    }

    .logo {
        margin-left: 50px;
    }

    .date-update {
        font-family: Open Sans, sans-serif;
        position: absolute;
        bottom: 0;
        right: 0;
        padding: 5px 0;
        font-size: 90%;
        width: 100%;
        text-align: end;
        padding-right: 30px;
        color: white;
        background-color: none;
        transition: 0.4s;
    }

    .date-update-mainstyle {
        color: rgb(96, 96, 96);
        background-color: rgb(238, 236, 245);
    }

    .topbar-main-style {
        background-color: white;
        color: rgb(37, 0, 132);
    }

    .flex-topbar {
        display: flex;
        align-items: center;
    }

    .logo {
        position: absolute;
        font-family: Open Sans, sans-serif;
        left: 30px;
        cursor: pointer;
    }

    .topbar-links {
        display: flex;
        height: 100%;
        font-family: Open Sans, sans-serif;
        align-items: center;
        position: absolute;
        right: 20px;
    }

    .point-of-bar {
        padding: 0.5rem 1rem;
        cursor: pointer;
    }
</style>