<script setup lang="ts">
import { onMounted } from 'vue';
import { Button } from '@/components/ui/button'
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger
} from '@/components/ui/tooltip'
import { FileCode, FileText, StickyNote } from 'lucide-vue-next'
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
    <div class="py-5 px-2">
        <h1 class="text-2xl font-black">
            RFC {{route.params.number}} | {{ title }}
        </h1>
        <p>
            {{ more_info }}
        </p>
      <div class="flex gap-2 mt-5">
        <div v-for="file in files" :key="file">
          <template v-if="file === 'html'">
            <TooltipProvider>
                <Tooltip>
                    <TooltipTrigger as-child>
                        <a :href="`https://www.rfc-editor.org/rfc/rfc${route.params.number}.html`">
                            <Button>
                                <FileCode />
                            </Button>
                        </a>
                    </TooltipTrigger>
                    <TooltipContent>
                        <p>HTML</p>
                    </TooltipContent>
                </Tooltip>
            </TooltipProvider>
          </template>
          <template v-else-if="file === 'pdf'">
                <TooltipProvider>
                    <Tooltip>
                        <TooltipTrigger as-child>
                            <a :href="`https://www.rfc-editor.org/rfc/pdfrfc/rfc${route.params.number}.txt.pdf`">
                                <Button >
                                    <StickyNote />
                                </Button>
                            </a>
                    </TooltipTrigger>
                <TooltipContent>
                        <p>PDF</p>
                </TooltipContent>
            </Tooltip>
            </TooltipProvider>
          </template>
          <template v-else-if="file === 'txt'">
            <TooltipProvider>
                    <Tooltip>
                        <TooltipTrigger as-child>
                            <a :href="`https://www.rfc-editor.org/rfc/rfc${route.params.number}.txt`">
                                <Button>
                                    <FileText />
                                </Button>
                            </a>
                        </TooltipTrigger>
                        <TooltipContent>
                            <p>TXT</p>
                        </TooltipContent>
                    </Tooltip>
                </TooltipProvider>
          </template>
        </div>
      </div>
    </div>
  </template>
  
