package allPaths

import (
    "path/filepath"
    "io/ioutil"
    "strings"
    "errors"
    "os"
)

var (
  extString string

  ignoreSlice []string
  paths []string

  justDirs bool = false
  justAll bool = true
)

func All( path string ) ( []string, error ) {

    if err := notEmptyString(path); err != nil {
        return paths, err
    }

    absoulte_path, err := filepath.Abs(path)
    if err != nil {
        return paths, err
    }

    if !filepath.IsAbs(absoulte_path) {
        return paths, errors.New("The path given is not a absolute path")
    }

    fs, err := ioutil.ReadDir(absoulte_path)
    if err != nil {
        return paths, err
    }
    if err := recursiveTreeDive(absoulte_path, fs); err != nil {
        return paths, err
    }
    return paths, nil
}

func Dirs( path string ) ( []string, error ) {
    if err := notEmptyString(path); err != nil {
        return paths, err
    }

    absoulte_path, err := filepath.Abs(path)
    if err != nil {
        return paths, err
    }

    if !filepath.IsAbs(absoulte_path) {
        return paths, errors.New("The path given is not a absolute path")
    }

    fs, err := ioutil.ReadDir(absoulte_path)
    if err != nil {
        return paths, err
    }

    justDirs = true
    justAll = false
    if err := recursiveTreeDive(absoulte_path, fs); err != nil {
        return paths, err
    }
    return paths, nil
}

func WithExt( path string, ext string ) ( []string, error ) {
      if err := notEmptyString(path); err != nil {
          return paths, err
      }

      absoulte_path, err := filepath.Abs(path)
      if err != nil {
          return paths, err
      }

      if !filepath.IsAbs(absoulte_path) {
          return paths, errors.New("The path given is not a absolute path")
      }

      fs, err := ioutil.ReadDir(absoulte_path)
      if err != nil {
          return paths, err
      }

      if !strings.Contains(ext, ".") {
          extString = "." + ext
          if err := recursiveTreeDive(absoulte_path, fs); err != nil {
            return paths, err
          }
          return paths, nil
      }
      if err := recursiveTreeDive(absoulte_path, fs); err != nil {
        return paths, err
      }
      return paths, nil
}

func notEmptyString( str string ) error {
	if len( str ) == 0 {
		  return errors.New("No data given to start path recursion")
	}
	return nil
}

// recursively gets all files paths in a given directory
func recursiveTreeDive( root string, fs []os.FileInfo ) error {

    for _, f := range fs {

        /*
            gets the file name, concatenate them with the given root,
            then cleans the path slash.
        */
        f := filepath.Clean( filepath.Join( root, f.Name() ) )

        fstat, err := os.Stat( f )
        if err != nil {
            return err
        }

        // checks weather the file is a directory
        if fstat.IsDir() {

            if justDirs {
               paths = append(paths, f )
            }

            newRoot := f

            fs, err := ioutil.ReadDir( newRoot )
            if err != nil {
                return err
            }

            // calls the function agian passing in the 'newRoot'
            // and ist directory contents
            if err := recursiveTreeDive( newRoot, fs ); err != nil {
                return err
            }
        } else {
            if justAll {

                base := filepath.Base(f)
                if extString != "" && filepath.Ext(base) == extString && filepath.Ext(base) != "" {
                    // appends the all file types name to the 'paths' slice array
                    paths = append(paths, f )
                } else if extString == "" {
                    // appends the all file types name to the 'paths' slice array
                    paths = append(paths, f )
                }
            }
        }
    }
    return nil
}

//
//
// if givenPath != "" {
//
//         absoulte_path, err := filepath.Abs(givenPath)
//         if err != nil {
//             fmt.Println(err)
//         } else {
//             givenPath := absoulte_path
//
//             if !fileExist("/./.cache-store") {
//                 createCacheStore()
//             } else if !initialize {
//                 cache_name = filepath.Base( givenPath ) + ".goc"
//
                // fs, err := ioutil.ReadDir( givenPath )
                // if err != nil {
                //   fmt.Println(err)
                // }

                // ignore()
                // recursiveTreeDive( givenPath, fs )
//
//                 strCachedPaths := strings.Join(cachedPaths, "\n")
//                 Err := writeStringToFile(filepath.Join( cacheStore, cache_name ), strCachedPaths)
//                 if Err != nil {
//                     fmt.Println(Err)
//                 }
//             } else {
//                 createCacheStore()
//             }
//         }
//     } else {
//         printError("run `cache --help` command to see how to use cache and what commands to use.")
//     }
// }

//
// func includes( str string, arr []string ) bool {
//     for _, i := range arr {
//         if str == i {
//             return true
//         }
//     }
//     return false
// }
//
// // these are files to gnore and not to cache
// func ignore() {
//
//     // checks if there is a cache-ignore file in the current directory
//     if fileExist("./cache-ignore") {
//         // gets the file name/paths from the file
//         data_stream, _ := readFile("./cache-ignore")
//         slice := strings.Split(data_stream, "\n")
//         ignoresSlice = slice[:len( slice ) - 1]
//     } else {
//
//         if len( ignoresSlice ) != 0 {
//             /*
//                 strips the squares brackets from the string
//                 splits the string by semicolons
//             */
//             ignoresSlice = strings.Split( strings.Trim(ignores, "[ ]"), "," )
//         } else {
//             ignoresSlice = []string { ".git", "node_modules" }
//         }
//     }
// }
//
