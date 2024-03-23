import { useEffect, useState } from 'react'
import './App.css'

interface Db {
  name: string;
  path: string;
}

interface TabelData {
  [key: string]: string;
}

interface Error {
  error: string;
  message: string;
}

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Toaster, toast } from 'sonner'
import MonacoEditor from './components/Editor';
import { Button } from './components/ui/button';
import { DataTable } from './components/sql-table/data-table';

function App() {
  const [dbs, setDbs] = useState<Db[]>([]);
  const [open, setOpen] = useState(true);
  const [value, setValue] = useState<string | undefined>('SELECT * FROM users;');
  const [tableData, setTableData] = useState<TabelData[]>([]);
  const [columns, setColumns] = useState<{ accessorKey: string, header: string }[]>([]);

  useEffect(() => {
    if (tableData.length > 0) {
      const firstRow = tableData[0];
      const newColumns = Object.keys(firstRow).map((key) => ({
        accessorKey: key,
        header: key,
      }));
      setColumns(newColumns);
    }
  }, [tableData]);

  useEffect(() => {
    fetch('/init')
      .then((res) => res.json())
      .then((data) => {
        setDbs(data);
      });
    setOpen(true);
  }, []);

  const handleDbClick = (db: Db) => {
    fetch('/connect', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(db),
    })
      .then((res) => {
        if (res.ok) {
          toast.success('Successfully connected to database');
          console.log('Successfully connected to database:', db.name);
          setOpen(false);
        } else {
          console.error('Failed to connect to database:', db.name);
        }
      })
      .catch((error) => console.error('Error connecting to database:', error));
  };

  const handleQuery = async () => {
    if (!value) return;

    try {
      const response = await fetch('/query', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ query: value }),
      });

      if (!response.ok) {
        const errorData: Error = await response.json();
        toast.error(errorData.error, {
          description: errorData.message,
        })
        throw new Error(errorData.message);
      }

      const data = await response.json();
      console.log(data);
      setTableData(data);
      toast.success('Query executed successfully');
    } catch (error) {
      console.error('Error executing query:', error);
    }
  };

  return (
    <div className='h-screen w-screen'>
      <Dialog open={open}>
        <DialogContent className="gap-0 p-0 outline-none bg-card">
          <DialogHeader className="px-6 pb-4 pt-5 border-b">
            <DialogTitle>Connect to you Database</DialogTitle>
            <DialogDescription>
              These are all the databases we found in your current repository. Select one to start querying.
            </DialogDescription>
          </DialogHeader>
          <div className='flex flex-col mx-auto my-4 h-48 overflow-scroll'>
            {
              dbs.map((db) => (
                <div key={db.path} onClick={() => handleDbClick(db)} className='relative flex flex-col gap-1 hover:bg-muted rounded sm:rounded-lg p-4 cursor-pointer'>
                  <h1 className='text-md font-semibold'>{db.name}</h1>
                  <p className='text-xs font-mono text-muted-foreground overflow-auto'>{db.path}</p>
                </div>
              ))
            }
          </div>
        </DialogContent>
      </Dialog>
      {!open &&
        <>
          <div className='h-1/2 w-4/5 flex flex-col mx-auto shadow-md mt-4 p-1 border-4 border-primary rounded-md overflow-auto'>
            <MonacoEditor
              code={String(value)}
              onChange={(value) => setValue(value)}
            />
            <div className='w-full flex justify-end pr-2 pb-2'>
              <Button onClick={handleQuery}>Run</Button>
            </div>
          </div>
          <div className="container mx-auto py-10">
            <DataTable columns={columns} data={tableData} />
          </div>
        </>
      }
      <Toaster richColors />
    </div>
  )
}

export default App
