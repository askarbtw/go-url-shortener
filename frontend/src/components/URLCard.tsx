import { useState } from 'react';
import { Box, Button, Flex, HStack, IconButton, Text, Tooltip, useToast, Badge } from '@chakra-ui/react';
import { FiCopy, FiEdit, FiExternalLink, FiTrash2, FiBarChart2 } from 'react-icons/fi';
import { useNavigate } from 'react-router-dom';
import { URL, URLStats } from '../types/url';
import { copyToClipboard, formatDate } from '../utils/helpers';
import { getRedirectURL } from '../services/api';

interface URLCardProps {
  urlData: URL | URLStats;
  onDelete?: (shortCode: string) => void;
  showStats?: boolean;
}

const URLCard = ({ urlData, onDelete, showStats = false }: URLCardProps) => {
  const navigate = useNavigate();
  const toast = useToast();
  const [isDeleting, setIsDeleting] = useState(false);

  const redirectUrl = getRedirectURL(urlData.shortCode);

  const handleCopy = async () => {
    const success = await copyToClipboard(redirectUrl);
    if (success) {
      toast({
        title: 'URL copied to clipboard',
        status: 'success',
        duration: 2000,
        isClosable: true,
      });
    } else {
      toast({
        title: 'Failed to copy URL',
        status: 'error',
        duration: 2000,
        isClosable: true,
      });
    }
  };

  const handleEdit = () => {
    navigate(`/edit/${urlData.shortCode}`);
  };

  const handleViewStats = () => {
    navigate(`/stats/${urlData.shortCode}`);
  };

  const handleDelete = async () => {
    if (!onDelete) return;
    
    setIsDeleting(true);
    try {
      await onDelete(urlData.shortCode);
      toast({
        title: 'URL deleted successfully',
        status: 'success',
        duration: 2000,
        isClosable: true,
      });
    } catch (error) {
      toast({
        title: 'Error deleting URL',
        status: 'error',
        duration: 2000,
        isClosable: true,
      });
    } finally {
      setIsDeleting(false);
    }
  };

  // Check if the URL has access count stats
  const hasStats = 'accessCount' in urlData;

  return (
    <Box borderWidth="1px" borderRadius="lg" overflow="hidden" p={4} bg="white" shadow="sm">
      <Flex direction="column" gap={3}>
        <Text fontWeight="bold" isTruncated title={urlData.url}>
          Original: {urlData.url}
        </Text>
        
        <Flex alignItems="center">
          <Text fontWeight="semibold" mr={2}>
            Short URL: 
          </Text>
          <Text color="blue.500" mr={2}>
            {redirectUrl}
          </Text>
          <HStack spacing={2}>
            <Tooltip label="Copy to clipboard">
              <IconButton
                aria-label="Copy URL"
                icon={<FiCopy />}
                size="sm"
                onClick={handleCopy}
              />
            </Tooltip>
            <Tooltip label="Open URL">
              <IconButton
                as="a"
                href={redirectUrl}
                target="_blank"
                aria-label="Open URL"
                icon={<FiExternalLink />}
                size="sm"
              />
            </Tooltip>
          </HStack>
        </Flex>

        {(showStats && hasStats) && (
          <Flex alignItems="center">
            <Text>
              Access Count: <Text as="span" fontWeight="bold">{(urlData as URLStats).accessCount}</Text>
            </Text>
            {!onDelete && (
              <Badge ml={2} colorScheme="blue">Stats Page</Badge>
            )}
          </Flex>
        )}

        <Text fontSize="sm" color="gray.500">
          Created: {formatDate(urlData.createdAt)}
        </Text>
        
        <Text fontSize="sm" color="gray.500">
          Last Updated: {formatDate(urlData.updatedAt)}
        </Text>

        {onDelete && (
          <Flex mt={2} justifyContent="flex-end">
            <Button
              size="sm"
              colorScheme="blue"
              variant="outline"
              leftIcon={<FiBarChart2 />}
              mr={2}
              onClick={handleViewStats}
            >
              Stats
            </Button>
            <Button
              size="sm"
              colorScheme="blue"
              variant="outline"
              leftIcon={<FiEdit />}
              mr={2}
              onClick={handleEdit}
            >
              Edit
            </Button>
            <Button
              size="sm"
              colorScheme="red"
              variant="outline"
              leftIcon={<FiTrash2 />}
              onClick={handleDelete}
              isLoading={isDeleting}
            >
              Delete
            </Button>
          </Flex>
        )}
      </Flex>
    </Box>
  );
};

export default URLCard; 