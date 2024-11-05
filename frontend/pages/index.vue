<script setup lang="ts">
import { ref } from 'vue'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Search, Eye} from 'lucide-vue-next'

import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'

const searchText = ref('');
const search = ref([]);

async function handleKeyUp(event: KeyboardEvent) {
  if (event.key === 'Enter') {
    const response = await getSearchData();
    search.value = await response.data;
  }
}

function setData(data){
    localStorage.setItem('files', data.files);
    localStorage.setItem('more_info', data.more_info);
    localStorage.setItem('title', data.title);
}

async function getSearchData() {
  return await $fetch(`http://localhost:4000/api/search?title=${searchText.value}`)
}
</script>

<template>
    <div class="bg-gray-900 h-[50vh]">
    <section>
        <div class="relative">
            <div class="absolute m-auto blur-[160px] max-w-s h-[13rem] top-12 inset-0"
                style="background:linear-gradient(180deg, #ff0000 0%, rgba(22, 22, 24, 0.984375) 0.01%, rgba(237, 78, 80, 0.2) 100%)">
            </div>
            <div class="relative">
                <div class="custom-screen py-28 relative">
                    <div class="relative z-10 duration-1000 delay-150 opacity-1">
                        <div class="max-w-2xl mx-auto text-center">
                            <span class="text-5xl font-bold bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 bg-clip-text text-transparent">RFC Search</span>
                            <h2 class="text-gray-50 text-2xl font-bold sm:text-2xl mt-2">The RFC Series</h2>
                            <p class="mt-2 text-gray-300">The RFC Series (ISSN 2070-1721) contains technical and organizational documents about the Internet, including the specifications and policy documents produced by five streams: the Internet Engineering Task Force (IETF), the Internet Research Task Force (IRTF), the Internet Architecture Board (IAB), Independent Submissions, and Editorial.</p>
                        </div>
                        <div class="flex justify-center font-medium text-md">
                            <div class="relative flex justify-center w-full py-9">
                            <div class="relative w-3/4 max-w-lg">
                                <Input
                                id="search"
                                type="text"
                                placeholder="Search..."
                                v-model="searchText"
                                class="pl-10 text-gray-100 bg-gray-900 border-solid border-2 border-gray-950 w-full"
                                @keyup="handleKeyUp"
                                />
                                <span class="absolute start-0 inset-y-0 flex items-center justify-center px-2">
                                <Search class="size-5 text-muted-foreground" />
                                </span>
                            </div>
                        </div>
                        </div>
                    </div>
                    <img src="https://i.ibb.co/8D7rcYv/download.webp" width="2880" height="1358" decoding="async" data-nimg="1" class="w-full h-full object-cover m-auto absolute inset-0 pointer-events-none" loading="lazy" style="color:transparent"/>
                </div>
            </div>
        </div>
    </section>
</div>

<Table class="px-5">
    <TableCaption>A list of RFC's</TableCaption>
    <TableHeader>
      <TableRow>
        <TableHead class="w-[100px]">Number</TableHead>
        <TableHead>Files</TableHead>
        <TableHead>Title</TableHead>
        <TableHead>Authors</TableHead>
        <TableHead>More Info</TableHead>
        <TableHead>Status</TableHead>
        <TableHead class="text-right">View</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow 
      v-for="item in search" 
      :key="item.number"
      style="cursor: pointer"
      href="/deneme"
      >
        <TableCell class="font-medium">
          {{ item.number }}
        </TableCell>
        <TableCell>{{ item.files.join(",") }}</TableCell>
        <TableCell>{{ item.title }}</TableCell>
        <TableCell>{{ item.authors.join(",") }}</TableCell>
        <TableCell>{{ item.more_info }}</TableCell>
        <TableCell class="text-right">
          {{ item.status }}
        </TableCell>
        <TableCell>
            <a :href="`rfc/${item.number}`">
                <Button @click="() => { 
                setData({
                    files: item.files.join(','),
                    title: item.title,
                    more_info: item.more_info
                }); 
                router.push(`/rfc/${item.number}`);
                }">
                    <Eye class="w-4 h-4 mr-2" />View 
                </Button>
            </a>
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>
  
