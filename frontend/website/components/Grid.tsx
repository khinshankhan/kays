import {
  Box,
  Image,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalHeader,
  ModalOverlay,
  Text
} from "@chakra-ui/react";
import { NextPage } from "next";
import React, { useState } from "react";
import { FileMeta } from "../util/file";

const emptyFileMeta = { uuid: "", filename: "" };

const FileGrid: NextPage<{ files: FileMeta[] }> = ({ files }) => {
  const [selectedFile, setSelectedFile] = useState<FileMeta>(emptyFileMeta);
  return (
    <Box display="flex" flexWrap="wrap" justifyContent="space-evenly">
      {files.map(file => (
        <Box
          key={file.uuid}
          p={1}
          border="1px solid black"
          cursor="pointer"
          onClick={() => setSelectedFile(file)}>
          <Box m={1} w={200} h={200} textAlign="center">
            <Image src={`/serve/${file.uuid}`} />
          </Box>
          <Text>
            {file.filename}
          </Text>
        </Box>
      ))}

      <Modal
        isOpen={selectedFile !== emptyFileMeta}
        onClose={() => setSelectedFile(emptyFileMeta)}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>{selectedFile.filename}</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Image src={`/serve/${selectedFile.uuid}`} />
            <Text>UUID: {selectedFile.uuid}</Text>
          </ModalBody>
        </ModalContent>
      </Modal>
    </Box>
  );
};

export default FileGrid;
