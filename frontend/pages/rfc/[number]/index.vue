<script setup lang="ts">
import { onMounted } from 'vue';
import { Button } from '@/components/ui/button'
import { FileCode, FileCode2, FileText, StickyNote } from 'lucide-vue-next'
const files = ref([]);
const title = ref("");
const more_info = ref("");
const route = useRoute();

onMounted(() => {
  const storedFiles = localStorage.getItem('files');
  const storedTitle = localStorage.getItem('title');
  const storedInfo = localStorage.getItem('more_info');
  if (storedFiles) {
    files.value = storedFiles.split(',');
    title.value = storedTitle;
    more_info.value = storedInfo;
  }
});
</script>

<template>
    <div>
        RFC {{route.params.number}}
        {{ title }}
        {{ more_info }}
      <div class="flex space-x-2">
        <div v-for="file in files" :key="file">
          <template v-if="file === 'html'">
            <a :href="`https://www.rfc-editor.org/rfc/rfc${route.params.number}.html`">
                <Button>
                    <FileCode />
                </Button>
            </a>
          </template>
          <template v-else-if="file === 'pdf'">
            <a :href="`https://www.rfc-editor.org/rfc/pdfrfc/rfc${route.params.number}.txt.pdf`">
                <Button >
                    <StickyNote />
                </Button>
            </a>
          </template>
          <template v-else-if="file === 'txt'">
            <a :href="`https://www.rfc-editor.org/rfc/rfc${route.params.number}.txt`">
                <Button>
                    <FileText />
                </Button>
            </a>
          </template>
        </div>
      </div>
    </div>
  </template>
  