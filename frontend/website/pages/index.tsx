import { Heading } from "@chakra-ui/react";
import { NextPage } from "next";
import Head from "next/head";
import { useEffect, useState } from "react";
import FileGrid from "../components/Grid";
import { FileMeta } from "../util/file";

const Index: NextPage = () => {
  const [files, setFiles] = useState<FileMeta[]>([]);

  useEffect(() => {
    fetch('/files')
      .then(res => res.json())
      .then(json => {
        setFiles(json);
      })
  }, []);

  return (
    <>
      <Head>
        <title>Caputo</title>
      </Head>

      <Heading as="h1" size="2xl" mb="2">
        Caputo
      </Heading>

      <FileGrid files={files} />
    </>
  );
};

export default Index;
