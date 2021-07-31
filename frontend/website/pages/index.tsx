import { Heading } from "@chakra-ui/react";
import { NextPage } from "next";
import Head from "next/head";

const Index: NextPage = () => {
  return (
    <>
      <Head>
        <title>Caputo</title>
      </Head>

      <Heading as="h1" size="2xl" mb="2">
        Hello World!
      </Heading>
    </>
  );
};

export default Index;
