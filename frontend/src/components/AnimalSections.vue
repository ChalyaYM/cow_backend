<template>
<div class="sections">
    <ul class="animal-links">
        <li @click="changeSection('common')" :class="{'selected': this.section.common}">Общая информация</li>
        <li @click="changeSection('rating')" :class="{'selected': this.section.rating}">Оценки</li>
        <li @click="changeSection('genetic')" :class="{'selected': this.section.genetic}">Генетика</li>
        <li @click="changeSection('exterior')" :class="{'selected': this.section.exterior}">Экстерьер</li>
        <li @click="changeSection('health')" :class="{'selected': this.section.health}">Ветеринарные мероприятия и события</li>
        <li @click="changeSection('milking')" :class="{'selected': this.section.milking}">Контрольные дойки и лактации</li>
        <li @click="changeSection('lineage')" :class="{'selected': this.section.lineage}">Родословная</li>
        <li @click="changeSection('matching')" :class="{'selected': this.section.matching}">Подбор пар</li>
        <li @click="changeSection('documents')" :class="{'selected': this.section.documents}">Документы</li> 
    </ul>
</div>
</template>

<script>
export default {
    data() {
        return {
            section: {
                common: true,
                rating: false,
                genetic: false,
                exterior: false,
                health: false,
                milking: false,
                lineage: false,
                matching: false,
                documents: false,
            }
        }
    },
    methods: {
        changeSection(section_name) {
            this.$emit('changeSection', section_name);
            for (let key in this.section) {
                this.section[key] = false;
                if (section_name == key) {
                    this.section[key] = true;
                }
            }
        }
    }
}
</script>

<style scoped>
.sections {
    font-family: Open Sans, sans-serif;
    position: sticky;
    top: 130px;
}

.animal-links {
    font-size: 110%;
    color:rgb(97, 49, 216);
    list-style: none;
}

.animal-links li {
    padding: 10px 5px;
    width: max-content;
    cursor: pointer;
    transition: 0.3s;
}

.animal-links li:hover {
    color:rgb(21, 6, 61);
}

.selected {
    color: rgb(21, 6, 61);
    font-weight: 520;
}
</style>